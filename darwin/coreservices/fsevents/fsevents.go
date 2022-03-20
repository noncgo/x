//go:build darwin
// +build darwin

// Package fsevents provides bindings for File System Events framework APIs.
//
// Overview
//
// The file system events API provides a way for your application to ask for
// notification when the contents of a directory hierarchy are modified. For
// example, your application can use this to quickly detect when the user
// modifies a file within a project bundle using another application.
//
// It also provides a lightweight way to determine whether the contents of a
// directory hierarchy have changed since your application last examined them.
// For example, a backup application can use this to determine what files have
// changed since a given time stamp or a given event ID.
//
// References
//  • https://developer.apple.com/documentation/coreservices/file_system_events
//  • https://developer.apple.com/library/archive/documentation/Darwin/Conceptual/FSEvents_ProgGuide/UsingtheFSEventsFramework/UsingtheFSEventsFramework.html
//  • https://developer.apple.com/documentation/coreservices
package fsevents
