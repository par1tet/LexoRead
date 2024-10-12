import cl from './../MainPage.module.css'

export function SearchPanel(){
    return (
        <div className={cl['searchPanel']}>
            <input placeholder='Search books' />
        </div>
    )
}
