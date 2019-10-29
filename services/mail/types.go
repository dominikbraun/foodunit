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

// Settings provides basic configuration values for sending an e-mail.
// It also holds placeholders which will be replaced with the mapped values.
type Settings struct {
	From      string            `json:"from"`
	To        string            `json:"to"`
	ToName    string            `json:"to_name"`
	Subject   string            `json:"subject"`
	Body      string            `json:"body"`
	Variables map[string]string `json:"variables"`
}
