package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type FileType string

const (
	FileTypeDirectory FileType = "directory"
	FileTypeFile      FileType = "file"
)

func main() {
	bytes, err := os.ReadFile("./input.txt")
	handleError(err)
	root := &File{Type: FileTypeDirectory}
	resolver := &Resolver{Current: root}
	for _, line := range strings.Split(string(bytes), "\n") {
		resolver.ProcessLine(line)
	}
	sum := calculateSumForAtMost(root, 100000)
	fmt.Printf("Sum: %d\n", sum)
	diskSize := 70000000
	required := 30000000
	missing := math.Abs(float64(diskSize - root.Size - required))
	min := findClosestSize(root, missing)
	fmt.Printf("Size of the directory to delete: %v (missing %v)\n", min, missing)
}

type File struct {
	Type     FileType
	Name     string
	Size     int
	Children []*File
	Parent   *File
}

func (f *File) AddToDirSize(size int) {
	f.Size += size
	if f.Parent != nil {
		f.Parent.AddToDirSize(size)
	}
}

type Resolver struct {
	Current *File
}

func (r *Resolver) ProcessLine(line string) {
	parts := strings.Split(line, " ")
	switch parts[0] {
	case "$":
		switch parts[1] {
		case "cd":
			r.GoTo(parts[2])
		default:
			return
		}
	default:
		r.ProcessFile(parts)
	}
}

func (r *Resolver) GoTo(path string) {
	if path == ".." && r.Current.Parent != nil {
		r.Current = r.Current.Parent
		return
	}

	for _, file := range r.Current.Children {
		if file.Name == path {
			r.Current = file
			return
		}
	}
}

func (r *Resolver) ProcessFile(parts []string) {
	child := NewFileFromLsOutput(parts, r.Current)
	r.Current.Children = append(r.Current.Children, child)
	if child.Type == FileTypeFile {
		r.Current.AddToDirSize(child.Size)
	}
}

func NewFileFromLsOutput(parts []string, parent *File) *File {
	if parts[0] == "dir" {
		return &File{
			Type:   FileTypeDirectory,
			Name:   parts[1],
			Parent: parent,
		}
	}

	size, err := strconv.Atoi(parts[0])
	handleError(err)
	return &File{
		Type:   FileTypeFile,
		Size:   size,
		Name:   parts[1],
		Parent: parent,
	}
}

func findClosestSize(dir *File, missing float64) float64 {
	min := math.MaxFloat64
	if float64(dir.Size) >= missing {
		min = math.Min(min, float64(dir.Size))
	}
	for _, child := range dir.Children {
		if child.Type == FileTypeDirectory {
			min = math.Min(min, findClosestSize(child, missing))
		}
	}
	return min
}

func calculateSumForAtMost(dir *File, maxSize int) int {
	if dir.Type != FileTypeDirectory {
		return 0
	}
	sum := 0
	if dir.Size <= maxSize {
		sum += dir.Size
	}
	for _, child := range dir.Children {
		sum += calculateSumForAtMost(child, maxSize)
	}
	return sum
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("an error has occurred: %v", err)
	}
}
