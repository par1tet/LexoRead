import IBook from './../interfaces/IBook.ts'

export class Book implements IBook {
    constructor(coverUrl: string){
        this.coverUrl = coverUrl
    }
}
