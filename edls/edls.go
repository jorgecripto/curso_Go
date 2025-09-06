package main

import "time"

// file types
const (
	fileRegular int = iota
	fileDirectory
	fileExecutable
	fileCompress
	fileImage
	fileLink
)

//file extensions

const (
	exe string = ".exe"
	deb string = ".deb"
	zip string = ".zip"
	gz  string = ".gz"
	tar string = ".tar"
	rar string = ".rar"
	png string = ".png"
	jpg string = ".jpg"
	gif string = ".gif"
)

type file struct {
	name             string
	fileType         int
	isDir            bool
	isHidden         bool
	userName         string
	groupName        string
	size             int64
	modificationTime time.Time
	mode             string
}

type styleTypeFile struct {
	icon   string
	color  string
	symbol string
}

var mapStyleByTypeFile = map[int]styleTypeFile{

	fileRegular:    {icon: "📰"},
	fileDirectory:  {icon: "📁", color: "BLUE", symbol: "/"},
	fileExecutable: {icon: "🚀", color: "GREEN", symbol: "*"},
	fileCompress:   {icon: "🗃️", color: "RED"},
	fileImage:      {icon: "📸", color: "MAGENTA"},
	fileLink:       {icon: "⛓️", color: "CYAN"},
}
