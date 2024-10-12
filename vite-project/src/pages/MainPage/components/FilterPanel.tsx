import cl from './../MainPage.module.css'
import { Filter } from './UI/Filter.tsx'

export function FilterPanel(){
    return (
        <div className={cl['filterPanel']}>
            <div className={cl['filterPanel__filters']}>
                <span></span>
            </div>
        </div>
    )
}
