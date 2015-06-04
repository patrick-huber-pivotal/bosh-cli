package templatescompiler

import (
	"os"
	"path/filepath"

	biproperty "github.com/cloudfoundry/bosh-init/common/property"
	bosherr "github.com/cloudfoundry/bosh-init/internal/github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-init/internal/github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-init/internal/github.com/cloudfoundry/bosh-utils/system"
	bireljob "github.com/cloudfoundry/bosh-init/release/job"
	bierbrenderer "github.com/cloudfoundry/bosh-init/templatescompiler/erbrenderer"
)

type JobRenderer interface {
	Render(releaseJob bireljob.Job, jobProperties, globalProperties biproperty.Map, deploymentName string) (RenderedJob, error)
}

type jobRenderer struct {
	erbRenderer bierbrenderer.ERBRenderer
	fs          boshsys.FileSystem
	logger      boshlog.Logger
	logTag      string
}

func NewJobRenderer(
	erbRenderer bierbrenderer.ERBRenderer,
	fs boshsys.FileSystem,
	logger boshlog.Logger,
) JobRenderer {
	return &jobRenderer{
		erbRenderer: erbRenderer,
		fs:          fs,
		logger:      logger,
		logTag:      "jobRenderer",
	}
}

func (r *jobRenderer) Render(releaseJob bireljob.Job, jobProperties, globalProperties biproperty.Map, deploymentName string) (RenderedJob, error) {
	context := NewJobEvaluationContext(releaseJob, jobProperties, globalProperties, deploymentName, r.logger)

	sourcePath := releaseJob.ExtractedPath

	destinationPath, err := r.fs.TempDir("rendered-jobs")
	if err != nil {
		return nil, bosherr.WrapError(err, "Creating rendered job directory")
	}

	renderedJob := NewRenderedJob(releaseJob, destinationPath, r.fs, r.logger)

	for src, dst := range releaseJob.Templates {
		err := r.renderFile(
			filepath.Join(sourcePath, "templates", src),
			filepath.Join(destinationPath, dst),
			context,
		)
		if err != nil {
			defer renderedJob.DeleteSilently()
			return nil, bosherr.WrapErrorf(err, "Rendering template src: %s, dst: %s", src, dst)
		}
	}

	err = r.renderFile(
		filepath.Join(sourcePath, "monit"),
		filepath.Join(destinationPath, "monit"),
		context,
	)
	if err != nil {
		defer renderedJob.DeleteSilently()
		return nil, bosherr.WrapError(err, "Rendering monit file")
	}

	return renderedJob, nil
}

func (r *jobRenderer) renderFile(sourcePath, destinationPath string, context bierbrenderer.TemplateEvaluationContext) error {
	err := r.fs.MkdirAll(filepath.Dir(destinationPath), os.ModePerm)
	if err != nil {
		return bosherr.WrapErrorf(err, "Creating tempdir '%s'", filepath.Dir(destinationPath))
	}

	err = r.erbRenderer.Render(sourcePath, destinationPath, context)
	if err != nil {
		return bosherr.WrapErrorf(err, "Rendering template src: %s, dst: %s", sourcePath, destinationPath)
	}
	return nil
}
