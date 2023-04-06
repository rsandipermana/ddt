package service

import (
    "strings"
)

type CountService interface {
    CountOccurrences(target string, text string) int
}

type countService struct {}

func NewCountService() CountService {
    return &countService{}
}

func (s *countService) CountOccurrences(target string, text string) int {
    return strings.Count(text, target)
}
