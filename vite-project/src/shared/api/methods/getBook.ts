import axios from 'axios'
import { bookService } from './../serviceLinks.ts'

export async function getBook(id: number) {
    console.log(id)
    let returnData = {}

    await axios.get(bookService(`books/id/${id}`))
    .then(r => (returnData = r.data))

    return returnData
}
