import cl from './BookPage.module.css'
import { Header } from './../../widgets/pageComponents/Header.tsx' 
import { useSearchParams } from 'react-router-dom'
import { getBook } from './../../shared/api/methods/getBook.ts'
import { useState, useEffect } from 'react'

export function BookPage() {
    const [searchParams, setSearchParams] = useSearchParams()
    const [bookData, setBookData] = useState({})
    console.log(bookData)

    useEffect(() => {
        getBook(searchParams.get("id"))
        .then(r => setBookData(r))
    }, [])

    if(bookData){
        return (
         <>
             <Header />
             <main className={cl['main']}>
                <div className={cl['main__bookInfo']}>
                    <div className={cl['main__bookInfo-cover']}>
                        <img src='https://i.pinimg.com/564x/9f/8b/37/9f8b377f90c0919cffb31a83d0eb8f36.jpg' />
                    </div>
                    <div className={cl['main__bookInfo-ohterInfo']}>
                        <div className={cl['main__bookInfo-author']}>
                            <span>{bookData.author}</span>
                        </div>
                        <div className={cl['main__bookInfo-title']}>
                            <span>“{bookData.title}”</span>
                        </div>
                        <div className={cl['main__bookInfo-read']}>
                            <button>
                                Читать
                            </button>
                        </div>
                        <div className={cl['main__bookInfo-desc']}>
                            <span>Описание</span>
                            <span>“{bookData.description}”</span>
                        </div>
                    </div>
                </div>
             </main>
         </>
        )
    }else{
        return (
            <>
                <main className={cl['main']}>Загрузка...</main>
            </>
        )
    }
}
