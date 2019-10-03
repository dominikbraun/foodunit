// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package dish provides services and types for sending mails.
package mail

import (
	"fmt"
	"github.com/dominikbraun/foodunit/storage"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"strings"
)

const (
	FromName string = "FoodUnit"
)

type Service struct {
	sgAPIKey string
	users    storage.User
}

func New(sgAPIKey string, u storage.User) *Service {
	service := Service{
		sgAPIKey: sgAPIKey,
		users:    u,
	}
	return &service
}

func (s *Service) Send(settings *Settings) error {
	for v, r := range settings.Variables {
		placeholder := fmt.Sprintf("{{%s}}", v)
		settings.Body = strings.Replace(settings.Body, placeholder, r, -1)
	}

	from := mail.NewEmail(FromName, settings.From)
	to := mail.NewEmail(settings.ToName, settings.To)
	email := mail.NewSingleEmail(from, settings.Subject, to, settings.Body, settings.Body)

	client := sendgrid.NewSendClient(s.sgAPIKey)

	_, err := client.Send(email)
	return err
}
