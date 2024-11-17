import IBook from './../../entities/interfaces/IBook.ts'
import { createBookList } from './../../features/createBookList.ts'
import cl from './ListBooks.module.css'
import { getBook } from './../../shared/api/methods/getBook.ts'
import { useNavigate } from 'react-router-dom'
type listBooksProps = {
    coverPaths: string[]
    title: string
}

export function ListBooks({coverPaths, title}: listBooksProps){
    const navigate = useNavigate()

    return (
        <div className={cl['listBooks']}> 
            <div className={cl['listBooks__title']}>
            
                <span>{title}</span>
            </div>
            <div className={cl['listBooks__books']}>
                {(createBookList(coverPaths)).map((book: IBook, index: number) =>
                    <img src={book.coverUrl} key={index} onClick={e => navigate("/book?id=2")}/>
                )}
            </div>
        </div>
    )
}
