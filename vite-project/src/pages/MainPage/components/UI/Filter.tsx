import cl from './../../BooksPage.module.css'
import clsx from 'clsx'

type FilterProps = {
    type: string
    placeholder: string
    className: string
}

export function Filter({type, placeholder, className}: FilterProps){
    return (
        <input type={type} placeholder={placeholder} className={clsx(className, cl['filter'])}/>
    )
}
