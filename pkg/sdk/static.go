package sdk

import "embed"

//go:embed static/*
var Static embed.FS

const StaticRootDirectory = "static"
const StaticConfigsDirectory = "static/configs"
const StaticMountDirectory = "static/mount"
