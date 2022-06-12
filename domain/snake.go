package domain

import (
	"errors"
	"math"
)

type Snake struct {
	squares                  []Square
	direction                Direction
	Frame                    uint64
	lastDirectionChangeFrame uint64
	remain                   uint64
	squareSize               int64
	originPosition           Position
}

func NewSnake(position Position, squareSize int64, snakeSize uint64) Snake {
	remain := snakeSize - 3
	squares := make([]Square, 0)

	if remain < 0 {
		remain = 0
	}

	for i := 0; i < 3; i++ {
		squares = append(squares, Square{
			Position: Position{
				X: position.X,
				Y: position.Y - int64(3-i+1)*int64(squareSize),
			},
			Size: squareSize,
		})
	}

	return Snake{
		squares:        squares,
		direction:      DirectionUp,
		squareSize:     squareSize,
		originPosition: position,
	}
}

func (s *Snake) getSquares() []Square {
	return s.squares
}

func (s *Snake) Move() error {
	switch s.direction {
	case DirectionUp:
		s.moveUp()
		break
	case DirectionLeft:
		s.moveLeft()
		break
	case DirectionRight:
		s.moveRight()
		break
	case DirectionDown:
		s.moveDown()
		break
	}

	for index, square := range s.squares {
		if s.IsTouching(square, &index) {
			return errors.New("Colission")
		}
	}

	if s.remain > 0 {
		s.Add(Square{
			Size: s.squareSize,
			Position: Position{
				X: s.originPosition.X,
				Y: s.originPosition.Y - 2*s.squareSize,
			},
		})

		s.remain--
	}

	s.Frame++

	return nil
}

func (s *Snake) size() int64 {
	return int64(len(s.squares) + int(s.remain))
}

func (s *Snake) moveUp() {
	var lastPosition Position

	for index, square := range s.squares {
		if index == 0 {
			lastPosition = square.Position
			s.squares[index] = Square{
				Size: square.Size,
				Position: Position{
					X: square.Position.X,
					Y: square.Position.Y - s.squareSize,
				},
			}
		} else {
			newSquare := Square{
				Size:     square.Size,
				Position: lastPosition,
			}

			lastPosition = square.Position

			s.squares[index] = newSquare
		}

	}
}

func (s *Snake) moveLeft() {
	var lastPosition Position

	for index, square := range s.squares {
		if index == 0 {
			lastPosition = square.Position
			s.squares[index] = Square{
				Size: square.Size,
				Position: Position{
					X: square.Position.X - s.squareSize,
					Y: square.Position.Y,
				},
			}
		} else {
			newSquare := Square{
				Size:     square.Size,
				Position: lastPosition,
			}

			lastPosition = square.Position

			s.squares[index] = newSquare
		}
	}
}

func (s *Snake) moveRight() {
	var lastPosition Position

	for index, square := range s.squares {
		if index == 0 {
			lastPosition = square.Position
			s.squares[index] = Square{
				Size: square.Size,
				Position: Position{
					X: square.Position.X + s.squareSize,
					Y: square.Position.Y,
				},
			}
		} else {
			newSquare := Square{
				Size:     square.Size,
				Position: lastPosition,
			}

			lastPosition = square.Position

			s.squares[index] = newSquare
		}
	}
}

func (s *Snake) moveDown() {
	var lastPosition Position

	for index, square := range s.squares {
		if index == 0 {
			lastPosition = square.Position
			s.squares[index] = Square{
				Size: square.Size,
				Position: Position{
					X: square.Position.X,
					Y: square.Position.Y + s.squareSize,
				},
			}
		} else {
			newSquare := Square{
				Size:     square.Size,
				Position: lastPosition,
			}

			lastPosition = square.Position

			s.squares[index] = newSquare
		}
	}
}

func (s *Snake) IsTouching(square Square, index *int) bool {
	isInsideSquare := func(position Position) bool {
		return math.Abs(float64(position.X-square.Position.X)) <= float64(square.Size/2) && math.Abs(float64(position.Y-square.Position.Y)) <= float64(square.Size/2)
	}

	if index != nil {
		for i, snakeSquares := range s.squares {
			if i != *index {
				if isInsideSquare(snakeSquares.Position) {
					return true
				}
			}
			return false
		}
	}

	for _, s := range s.squares {

		if isInsideSquare(s.Position) {
			return true
		}

	}

	return false
}

func (s *Snake) Add(square Square) {
	s.squares = append(s.squares, square)
}

func (s *Snake) SetDirection(direction Direction) {
	if s.Frame == s.lastDirectionChangeFrame {
		return
	}

	if direction == DirectionUp && s.direction == DirectionDown {
		return
	}

	if direction == DirectionDown && s.direction == DirectionUp {
		return
	}

	if direction == DirectionLeft && s.direction == DirectionRight {
		return
	}

	if direction == DirectionRight && s.direction == DirectionLeft {
		return
	}

	s.direction = direction
	s.lastDirectionChangeFrame = s.Frame
}

func (s *Snake) GetSquares() []Square {
	return s.squares
}
