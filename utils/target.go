package utils

type Target struct {
	Name string
	Host string
	Flag string
}

var EntireNodeTarget Target = Target{
	Name: "entire node",
	Host: "node",
	Flag: "all",
}

var NodeControllerTarget Target = Target{
	Name: "node controller",
	Host: "node",
	Flag: "nc",
}

var EdgeProcessorTarget Target = Target{
	Name: "edge processor",
	Host: "edge",
	Flag: "ep",
}

var WagmanTarget Target = Target{
	Name: "wagman",
	Host: "node",
	Flag: "wgn",
}
