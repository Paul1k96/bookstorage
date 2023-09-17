package book_grpc

import (
	"context"
	"github.com/Paul1k96/bookstorage/internal/models"
	"github.com/Paul1k96/bookstorage/internal/modules/books/service"
	mock_service "github.com/Paul1k96/bookstorage/internal/modules/books/service/mocks"
	"github.com/Paul1k96/bookstorage/pkg/grpc/books"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

type BooksServiceGRPCSuite struct {
	suite.Suite
	*require.Assertions

	ctrl             *gomock.Controller
	mockBooksService *mock_service.MockBookServicer

	client books.BookServiceGRPCClient
	closer func()
}

func TestBooksServiceGRPCSuite(t *testing.T) {
	suite.Run(t, new(BooksServiceGRPCSuite))
}

func (b *BooksServiceGRPCSuite) SetupTest() {
	b.Assertions = require.New(b.T())

	b.ctrl = gomock.NewController(b.T())
	b.mockBooksService = mock_service.NewMockBookServicer(b.ctrl)

	b.client, b.closer = b.getGRPC(context.Background())
}

func (b *BooksServiceGRPCSuite) TearDownTest() {
	defer b.closer()
	b.ctrl.Finish()
}

func (b *BooksServiceGRPCSuite) getGRPC(ctx context.Context) (books.BookServiceGRPCClient, func()) {
	buffer := 1024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	books.RegisterBookServiceGRPCServer(baseServer, NewBooksServiceGRPC(b.mockBooksService))
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := books.NewBookServiceGRPCClient(conn)

	return client, closer
}

func (b *BooksServiceGRPCSuite) TestBookService_GetBooksByAuthor() {
	serviceIn := service.GetBooksByAuthorIn{Name: "Author"}
	serviceOut := service.GetBooksByAuthorOut{Book: []models.Book{{ID: 1, Name: "Book One"}, {ID: 2, Name: "Book Two"}, {ID: 3, Name: "Book Three"}}, ErrorCode: 0}

	b.mockBooksService.EXPECT().GetBooksByAuthor(serviceIn).Return(serviceOut).Times(1)

	expectedGRPCOut := &books.GetBooksByAuthorOut{
		Books: []*books.Book{{Id: 1, Name: "Book One"}, {Id: 2, Name: "Book Two"}, {Id: 3, Name: "Book Three"}},
	}

	result, err := b.client.GetBooksByAuthor(context.Background(), &books.GetBooksByAuthorIn{Author: "Author"})

	b.Equal(err, nil)
	b.Equal(result.ErrorCode, int32(0))
	for i := range expectedGRPCOut.Books {
		b.Equal(expectedGRPCOut.Books[i], result.Books[i])
	}
}

func (b *BooksServiceGRPCSuite) TestBookService_GetAuthorsByBook() {
	serviceIn := service.GetAuthorsByBookIn{Name: "Book"}
	serviceOut := service.GetAuthorsByBookOut{Author: []models.Author{{ID: 1, Name: "Author One"}, {ID: 2, Name: "Author Two"}, {ID: 3, Name: "Author Three"}}, ErrorCode: 0}

	b.mockBooksService.EXPECT().GetAuthorsByBook(serviceIn).Return(serviceOut).Times(1)

	expectedGRPCOut := &books.GetAuthorsByBookOut{
		Authors: []*books.Author{{Id: 1, Name: "Author One"}, {Id: 2, Name: "Author Two"}, {Id: 3, Name: "Author Three"}},
	}

	result, err := b.client.GetAuthorsByBook(context.Background(), &books.GetAuthorsByBookIn{Book: "Book"})

	b.Equal(err, nil)
	b.Equal(result.ErrorCode, int32(0))
	for i := range expectedGRPCOut.Authors {
		b.Equal(expectedGRPCOut.Authors[i], result.Authors[i])
	}
}
