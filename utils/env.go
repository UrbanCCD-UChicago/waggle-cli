package utils

type Env struct {
	Name              string
	AccessibleTargets []Target
}

var AdminEnv Env = Env{
	Name: "admin",
	AccessibleTargets: []Target{
		EntireNodeTarget,
		NodeControllerTarget,
		EdgeProcessorTarget,
		WagmanTarget,
	},
}

var NodeControllerEnv Env = Env{
	Name: "node controller",
	AccessibleTargets: []Target{
		NodeControllerTarget,
		EdgeProcessorTarget,
		WagmanTarget,
	},
}

var EdgeProcessorEnv Env = Env{
	Name: "edge processor",
	AccessibleTargets: []Target{
		EdgeProcessorTarget,
	},
}
