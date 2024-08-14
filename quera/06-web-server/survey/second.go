package main

//import "errors"
//
//type Comment struct {
//	// do not modify or remove these fields
//	Score int
//	Text  string
//	// but you can add anything you want
//}
//
//type Survey struct {
//	flightRateSums  map[string]int
//	flightRateCount map[string]int
//	flightComments  map[string][]string
//	ticketFeedbacks map[[2]string]bool
//}
//
//func NewSurvey() *Survey {
//	return &Survey{
//		flightRateCount: make(map[string]int),
//		flightRateSums:  make(map[string]int),
//		flightComments:  make(map[string][]string),
//		ticketFeedbacks: make(map[[2]string]bool),
//	}
//}
//
//func (s *Survey) AddFlight(flightName string) error {
//	if _, found := s.flightRateCount[flightName]; found {
//		return errors.New("duplicate flight")
//	}
//
//	s.flightRateSums[flightName] = 0
//	s.flightRateCount[flightName] = 0
//	s.flightComments[flightName] = make([]string, 0)
//	return nil
//}
//
//func (s *Survey) AddTicket(flightName, passengerName string) error {
//	if _, found := s.flightRateCount[flightName]; !found {
//		return errors.New("flight not found")
//	}
//
//	if _, found := s.ticketFeedbacks[[2]string{flightName, passengerName}]; found {
//		return errors.New("duplicate ticket")
//	}
//
//	s.ticketFeedbacks[[2]string{flightName, passengerName}] = false
//	return nil
//}
//
//func (s *Survey) AddComment(flightName, passengerName string, comment Comment) error {
//	if _, found := s.flightRateCount[flightName]; !found {
//		return errors.New("flight not found")
//	}
//
//	hasFeedback, found := s.ticketFeedbacks[[2]string{flightName, passengerName}]
//	if !found {
//		return errors.New("ticket not found")
//	}
//
//	if hasFeedback {
//		return errors.New("feedback is already given")
//	}
//
//	if comment.Score < 1 || comment.Score > 10 {
//		return errors.New("invalid score")
//	}
//
//	s.ticketFeedbacks[[2]string{flightName, passengerName}] = true
//	s.flightRateSums[flightName] += comment.Score
//	s.flightRateCount[flightName]++
//	s.flightComments[flightName] = append(s.flightComments[flightName], comment.Text)
//	return nil
//}
//
//func (s *Survey) GetCommentsAverage(flightName string) (float64, error) {
//	if _, found := s.flightRateCount[flightName]; !found {
//		return 0, errors.New("flight not found")
//	}
//	if s.flightRateCount[flightName] == 0 {
//		return 0, errors.New("no comment submitted for this flight")
//	}
//	return float64(s.flightRateSums[flightName]) / float64(s.flightRateCount[flightName]), nil
//}
//
//func (s *Survey) GetAllCommentsAverage() map[string]float64 {
//	result := make(map[string]float64)
//	for flight, score := range s.flightRateSums {
//		if s.flightRateCount[flight] == 0 {
//			continue
//		}
//		result[flight] = float64(score) / float64(s.flightRateCount[flight])
//	}
//	return result
//}
//
//func (s *Survey) GetComments(flightName string) ([]string, error) {
//	if _, found := s.flightRateCount[flightName]; !found {
//		return nil, errors.New("flight not found")
//	}
//	return s.flightComments[flightName], nil
//}
//
//func (s *Survey) GetAllComments() map[string][]string {
//	return s.flightComments
//}
