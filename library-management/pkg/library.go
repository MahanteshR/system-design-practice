package pkg

import (
	"fmt"
)

type Library struct {
	Racks         []Rack
	BorrowedBooks []BorrowRecord
}

func (l *Library) AddBook(book Book, copies int) error {
	for i := range l.Racks {
		if l.Racks[i].Capacity-l.Racks[i].BookCount() >= copies {
			l.Racks[i].Books[book.ID] += copies

			return nil
		}
	}

	return fmt.Errorf("Not enough space to add %d copies of book %s", copies, book.Title)
}

func (l *Library) SearchBook(bookID string) (string, int, error) {
	for _, rack := range l.Racks {
		if cnt, found := rack.Books[bookID]; found {
			return rack.RackID, cnt, nil
		}
	}

	return "", 0, fmt.Errorf("bookID %s not found", bookID)
}

func (l *Library) RemoveBook(bookID string, copies int) error {
	for _, rack := range l.Racks {
		if cnt, found := rack.Books[bookID]; found && cnt >= copies {
			rack.Books[bookID] -= copies
			if rack.Books[bookID] == 0 {
				delete(rack.Books, bookID)
			}

			return nil
		}
	}

	return fmt.Errorf("Not enough copies of book %s to remove", bookID)
}

func (l *Library) BorrowBook(bookID, borrower, dueDate string) error {
	for _, rack := range l.Racks {
		if cnt, found := rack.Books[bookID]; found && cnt > 0 {
			rack.Books[bookID] -= 1
			if rack.Books[bookID] == 0 {
				delete(rack.Books, bookID)
			}

			l.BorrowedBooks = append(l.BorrowedBooks, BorrowRecord{
				BookID:   bookID,
				DueDate:  dueDate,
				Borrower: borrower,
			})

			return nil
		}
	}

	return fmt.Errorf("Book %s is not available for borrowing", bookID)
}
