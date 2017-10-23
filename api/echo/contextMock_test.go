package echo

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
	"net/url"
	"mime/multipart"
	"io"
	"net/http"
)

type contextMock struct {
	mock.Mock
}

func (c *contextMock) Request() *http.Request {
	panic("not implemented")
}

func (c *contextMock) SetRequest(r *http.Request) {
	panic("not implemented")
}

func (c *contextMock) Response() *echo.Response {
	panic("not implemented")
}

func (c *contextMock) IsTLS() bool {
	panic("not implemented")
}

func (c *contextMock) IsWebSocket() bool {
	panic("not implemented")
}

func (c *contextMock) Scheme() string {
	panic("not implemented")
}

func (c *contextMock) RealIP() string {
	panic("not implemented")
}

func (c *contextMock) Path() string {
	panic("not implemented")
}

func (c *contextMock) SetPath(p string) {
	panic("not implemented")
}

func (c *contextMock) Param(name string) string {
	panic("not implemented")
}

func (c *contextMock) ParamNames() []string {
	panic("not implemented")
}

func (c *contextMock) SetParamNames(names ...string) {
	panic("not implemented")
}

func (c *contextMock) ParamValues() []string {
	panic("not implemented")
}

func (c *contextMock) SetParamValues(values ...string) {
	panic("not implemented")
}

func (c *contextMock) QueryParam(name string) string {
	panic("not implemented")
}

func (c *contextMock) QueryParams() url.Values {
	panic("not implemented")
}

func (c *contextMock) QueryString() string {
	panic("not implemented")
}

func (c *contextMock) FormValue(name string) string {
	panic("not implemented")
}

func (c *contextMock) FormParams() (url.Values, error) {
	panic("not implemented")
}

func (c *contextMock) FormFile(name string) (*multipart.FileHeader, error) {
	panic("not implemented")
}

func (c *contextMock) MultipartForm() (*multipart.Form, error) {
	panic("not implemented")
}

func (c *contextMock) Cookie(name string) (*http.Cookie, error) {
	panic("not implemented")
}

func (c *contextMock) SetCookie(cookie *http.Cookie) {
	panic("not implemented")
}

func (c *contextMock) Cookies() []*http.Cookie {
	panic("not implemented")
}

func (c *contextMock) Get(key string) interface{} {
	panic("not implemented")
}

func (c *contextMock) Set(key string, val interface{}) {
	panic("not implemented")
}

func (c *contextMock) Bind(i interface{}) error {
	panic("not implemented")
}

func (c *contextMock) Validate(i interface{}) error {
	panic("not implemented")
}

func (c *contextMock) Render(code int, name string, data interface{}) error {
	panic("not implemented")
}

func (c *contextMock) HTML(code int, html string) error {
	panic("not implemented")
}

func (c *contextMock) HTMLBlob(code int, b []byte) error {
	panic("not implemented")
}

func (c *contextMock) String(code int, s string) error {
	panic("not implemented")
}

func (c *contextMock) JSON(code int, i interface{}) error {
	panic("not implemented")
}

func (c *contextMock) JSONPretty(code int, i interface{}, indent string) error {
	panic("not implemented")
}

func (c *contextMock) JSONBlob(code int, b []byte) error {
	panic("not implemented")
}

func (c *contextMock) JSONP(code int, callback string, i interface{}) error {
	panic("not implemented")
}

func (c *contextMock) JSONPBlob(code int, callback string, b []byte) error {
	panic("not implemented")
}

func (c *contextMock) XML(code int, i interface{}) error {
	panic("not implemented")
}

func (c *contextMock) XMLPretty(code int, i interface{}, indent string) error {
	panic("not implemented")
}

func (c *contextMock) XMLBlob(code int, b []byte) error {
	panic("not implemented")
}

func (c *contextMock) Blob(code int, contentType string, b []byte) error {
	panic("not implemented")
}

func (c *contextMock) Stream(code int, contentType string, r io.Reader) error {
	panic("not implemented")
}

func (c *contextMock) File(file string) error {
	panic("not implemented")
}

func (c *contextMock) Attachment(file string, name string) error {
	panic("not implemented")
}

func (c *contextMock) Inline(file string, name string) error {
	panic("not implemented")
}

func (c *contextMock) NoContent(code int) error {
	panic("not implemented")
}

func (c *contextMock) Redirect(code int, url string) error {
	panic("not implemented")
}

func (c *contextMock) Error(err error) {
	panic("not implemented")
}

func (c *contextMock) Handler() echo.HandlerFunc {
	panic("not implemented")
}

func (c *contextMock) SetHandler(h echo.HandlerFunc) {
	panic("not implemented")
}

func (c *contextMock) Logger() echo.Logger {
	panic("not implemented")
}

func (c *contextMock) Echo() *echo.Echo {
	panic("not implemented")
}

func (c *contextMock) Reset(r *http.Request, w http.ResponseWriter) {
	panic("not implemented")
}

