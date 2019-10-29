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

// Service executes mail-related business logic and use cases. It is also responsible
// for accessing the model storage under consideration of all business rules.
type Service struct {
	sgAPIKey string
	users    storage.User
}

// NewService creates a new Service instance utilizing the given storage objects.
// The storage objects need to be ready to use for the service.
func New(sgAPIKey string, u storage.User) *Service {
	service := Service{
		sgAPIKey: sgAPIKey,
		users:    u,
	}
	return &service
}

// Send creates an e-mail and attempts to send it via the SendGrid API. It will
// replace all placeholders of Settings.Variables in the mail subject and body.
//
// For instance, the subject may look like this: `Welcome to FoodUnit, {{name}}!`
// To replace {{name}} with a real value, you should provide settings like this:
//
//		settings := mail.Settings{
//			Subject: `Welcome to FoodUnit, {{name}}!`,
//			Variables: map[string]string{
//				"name": username,
//			}
//		}
//
//		mailService.Send(&settings)
//
// Send will automatically replace the map key with the associated value. Note that
// the key must not include the curly brackets.
func (s *Service) Send(settings *Settings) error {
	for variable, replacement := range settings.Variables {
		placeholder := fmt.Sprintf("{{%s}}", variable)
		settings.Subject = strings.Replace(settings.Subject, placeholder, replacement, -1)
		settings.Body = strings.Replace(settings.Body, placeholder, replacement, -1)
	}

	from := mail.NewEmail(FromName, settings.From)
	to := mail.NewEmail(settings.ToName, settings.To)
	email := mail.NewSingleEmail(from, settings.Subject, to, settings.Body, settings.Body)

	client := sendgrid.NewSendClient(s.sgAPIKey)

	_, err := client.Send(email)
	return err
}
