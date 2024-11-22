import axios from 'axios';
import { userService } from '../serviceLinks.ts';
export async function changeDescription(id: number | null) {
  console.log(id);
  let description = ''
  console.log(description);
  await axios.put(userService(`changeDescription/id/${id}`)).then((response) => {
    description += response.data.description;

  })

  return description
}