package service

import (
	"errors"
	"testing"

	serviceErr "github.com/Paul1k96/bookstorage/internal/infrastructure/errors"
	"github.com/Paul1k96/bookstorage/internal/models"
	mock_storage "github.com/Paul1k96/bookstorage/internal/modules/books/storage/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type BooksServiceSuite struct {
	suite.Suite
	*require.Assertions

	ctrl             *gomock.Controller
	mockBooksStorage *mock_storage.MockBookStorager

	service BookServicer
}

func TestBooksServiceSuite(t *testing.T) {
	suite.Run(t, new(BooksServiceSuite))
}

func (b *BooksServiceSuite) SetupTest() {
	b.Assertions = require.New(b.T())

	b.ctrl = gomock.NewController(b.T())
	b.mockBooksStorage = mock_storage.NewMockBookStorager(b.ctrl)

	b.service = NewBookService(b.mockBooksStorage)
}

func (b *BooksServiceSuite) TearDownTest() {
	b.ctrl.Finish()
}

func (b *BooksServiceSuite) TestBookService_GetBooksByAuthor() {
	author := "Author"
	storageOut := []models.Book{{ID: 1, Name: "Book One"}, {ID: 2, Name: "Book Two"}, {ID: 3, Name: "Book Three"}}

	b.mockBooksStorage.EXPECT().GetBooksByAuthor(author).Return(storageOut, nil).Times(1)

	expectedOut := GetBooksByAuthorOut{
		Book: storageOut,
	}

	result := b.service.GetBooksByAuthor(GetBooksByAuthorIn{
		Name: author,
	})

	b.Equal(expectedOut, result)
}

func (b *BooksServiceSuite) TestBookService_GetBooksByAuthorNegative() {
	err := errors.New("create err")
	author := "Author"

	b.mockBooksStorage.EXPECT().GetBooksByAuthor(author).Return(nil, err).Times(1)

	expectedOut := GetBooksByAuthorOut{
		ErrorCode: serviceErr.BookServiceGetBooksByAuthorErr,
	}

	result := b.service.GetBooksByAuthor(GetBooksByAuthorIn{
		Name: author,
	})

	b.Equal(expectedOut, result)

}

func (b *BooksServiceSuite) TestBookService_GetAuthorsByBook() {
	book := "Book"
	storageOut := []models.Author{{ID: 1, Name: "Author 1"}, {ID: 2, Name: "Author 2"}, {ID: 3, Name: "Author 3"}}

	b.mockBooksStorage.EXPECT().GetAuthorsByBook(book).Return(storageOut, nil).Times(1)

	expectedOut := GetAuthorsByBookOut{
		Author: storageOut,
	}

	result := b.service.GetAuthorsByBook(GetAuthorsByBookIn{
		Name: book,
	})

	b.Equal(expectedOut, result)
}

func (b *BooksServiceSuite) TestBookService_GetAuthorsByBookNegative() {
	err := errors.New("create error")
	book := "Book"

	b.mockBooksStorage.EXPECT().GetAuthorsByBook(book).Return(nil, err).Times(1)

	expectedOut := GetAuthorsByBookOut{
		ErrorCode: serviceErr.BookServiceGetAuthorsByBookErr,
	}

	result := b.service.GetAuthorsByBook(GetAuthorsByBookIn{
		Name: book,
	})

	b.Equal(expectedOut, result)
}
