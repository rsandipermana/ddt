package service

import "github.com/rsandipermana/ddt/model"

type MaxService interface {
    FindMax(numbers []int) int
}

type maxService struct {}

func NewMaxService() MaxService {
    return &maxService{}
}

func (s *maxService) FindMax(numbers []int) int {
    max := numbers[0]
    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    return max
}
