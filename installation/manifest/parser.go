package manifest

import (
	biproperty "github.com/cloudfoundry/bosh-init/common/property"
	bosherr "github.com/cloudfoundry/bosh-init/internal/github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-init/internal/github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-init/internal/github.com/cloudfoundry/bosh-utils/system"
	boshuuid "github.com/cloudfoundry/bosh-init/internal/github.com/cloudfoundry/bosh-utils/uuid"
	"github.com/cloudfoundry/bosh-init/internal/gopkg.in/yaml.v2"
	birelsetmanifest "github.com/cloudfoundry/bosh-init/release/set/manifest"
)

type Parser interface {
	Parse(path string, releaseSetManifest birelsetmanifest.Manifest) (Manifest, error)
}

type parser struct {
	fs            boshsys.FileSystem
	uuidGenerator boshuuid.Generator
	logger        boshlog.Logger
	logTag        string
	validator     Validator
}

type manifest struct {
	Name          string
	CloudProvider installation `yaml:"cloud_provider"`
}

type installation struct {
	Template   template
	Properties map[interface{}]interface{}
	SSHTunnel  SSHTunnel `yaml:"ssh_tunnel"`
	Mbus       string
}

func (i installation) HasSSHTunnel() bool {
	return i.SSHTunnel != SSHTunnel{}
}

type template struct {
	Name    string
	Release string
}

func NewParser(fs boshsys.FileSystem, uuidGenerator boshuuid.Generator, logger boshlog.Logger, validator Validator) Parser {
	return &parser{
		fs:            fs,
		uuidGenerator: uuidGenerator,
		logger:        logger,
		logTag:        "deploymentParser",
		validator:     validator,
	}
}

func (p *parser) Parse(path string, releaseSetManifest birelsetmanifest.Manifest) (Manifest, error) {
	contents, err := p.fs.ReadFile(path)
	if err != nil {
		return Manifest{}, bosherr.WrapErrorf(err, "Reading file %s", path)
	}

	comboManifest := manifest{}
	err = yaml.Unmarshal(contents, &comboManifest)
	if err != nil {
		return Manifest{}, bosherr.WrapError(err, "Unmarshalling installation manifest")
	}
	p.logger.Debug(p.logTag, "Parsed installation manifest: %#v", comboManifest)

	if comboManifest.CloudProvider.SSHTunnel.PrivateKey != "" {
		privateKeyPath, err := p.fs.ExpandPath(comboManifest.CloudProvider.SSHTunnel.PrivateKey)
		if err != nil {
			p.logger.Warn(p.logTag, "Failed to expand private key path, using original path")
			privateKeyPath = comboManifest.CloudProvider.SSHTunnel.PrivateKey
		}
		comboManifest.CloudProvider.SSHTunnel.PrivateKey = privateKeyPath
	}

	installationManifest := Manifest{
		Name: comboManifest.Name,
		Template: ReleaseJobRef{
			Name:    comboManifest.CloudProvider.Template.Name,
			Release: comboManifest.CloudProvider.Template.Release,
		},
		Mbus: comboManifest.CloudProvider.Mbus,
	}

	properties, err := biproperty.BuildMap(comboManifest.CloudProvider.Properties)
	if err != nil {
		return Manifest{}, bosherr.WrapErrorf(err, "Parsing cloud_provider manifest properties: %#v", comboManifest.CloudProvider.Properties)
	}
	installationManifest.Properties = properties

	if comboManifest.CloudProvider.HasSSHTunnel() {
		password, err := p.uuidGenerator.Generate()
		if err != nil {
			return Manifest{}, bosherr.WrapError(err, "Generating registry password")
		}
		installationManifest.PopulateRegistry("registry", password, "127.0.0.1", 6901, comboManifest.CloudProvider.SSHTunnel)
	}

	err = p.validator.Validate(installationManifest, releaseSetManifest)
	if err != nil {
		return Manifest{}, bosherr.WrapError(err, "Validating installation manifest")
	}

	return installationManifest, nil
}
