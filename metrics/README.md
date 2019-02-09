# metrics
Metrics

## Usage

```
import "metrics"

var met metrics.Metrics

func timerExampele() {
	met.Timer.Start()

	//stuff to time

	met.Timer.Stop()
	met.Update(true, false)
}

```
```
func counterExample() {

	//stuff to count

	met.Counter.Increment()
	met.Update(false, true)
}
```

```
...
...
met.Publish()

```

