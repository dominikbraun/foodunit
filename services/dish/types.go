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

// Package dish provides services and types for Dish-related data.
package dish

// Variant is the API output for model.Variant.
type Variant struct {
	ID        uint64 `jsom:"id"`
	Value     string `json:"value"`
	IsDefault bool   `json:"is_default"`
	Price     uint   `json:"price"`
}

// Characteristic is the API output for model.Characteristic.
type Characteristic struct {
	ID       uint64    `json:"id"`
	Name     string    `json:"name"`
	Multiple bool      `json:"multiple"`
	Variants []Variant `json:"variants"`
}
