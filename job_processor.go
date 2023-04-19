package jobprocessor

type Job map[string]interface{}

type JobProcessor interface {
	Process(Job) error
}

type JobType string

type ExecutionContext map[string]interface{}

type Factory func(ctx ExecutionContext) (JobProcessor, error)

type Registry map[JobType]Factory

var jobProcessorRegistry Registry = map[JobType]Factory{}

func RegisterFactoryForJobType(jobType JobType, jobProcFactory Factory) {
	jobProcessorRegistry[jobType] = jobProcFactory
}

func GetFactoryForJobType(jobType JobType) Factory {
	return jobProcessorRegistry[jobType]
}
