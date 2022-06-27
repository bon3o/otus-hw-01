package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := doneStage(in, done)
	for _, stage := range stages {
		if stage == nil {
			panic("stage cant be nil")
		}
		out = stage(doneStage(out, done))
	}
	return out
}

func doneStage(in In, done In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case value, ok := <-in:
				if !ok {
					return
				}
				out <- value
			}
		}
	}()
	return out
}
