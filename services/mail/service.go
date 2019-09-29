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

import "github.com/dominikbraun/foodunit/storage"

type Config struct {
	From    string `json:"from"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type Service struct {
	config   *Config
	sgAPIKey string
	users    storage.User
}

func New(config *Config, sgAPIKey string, u storage.User) *Service {
	service := Service{
		config:   config,
		sgAPIKey: sgAPIKey,
		users:    u,
	}
	return &service
}