import { Book } from './../entities/classes/book.ts'

export function createBookList(coverUrls: string[]): Book[]{
    const booksArrays: Book[] = []

    coverUrls.forEach(cover => booksArrays.push(new Book(cover)))

    return booksArrays
}
