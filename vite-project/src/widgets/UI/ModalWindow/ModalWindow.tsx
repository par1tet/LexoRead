import cl from './ModalWindow.module.css';
import { NavLink } from 'react-router-dom';
import { jwtDecode } from 'jwt-decode';

export default function ModalWindow() {
  interface JwtPayload {
    id: number;
  }

  localStorage.setItem('jwtToken', "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJzbWFsbGdpZ2FjaCIsImVtYWlsIjoic21hbGxnaWdhY2hAbGV4b3JlYWQucnUiLCJoYXNoUGFzc3dvcmQiOiIkMmIkMDYkS01mdFY1c3kxRzZGQkpaa0dadnhqT0NBZEN4eVM2UFFJMFl2WDlqTmVrTEpGeHdBZmRaYnUifQ.9LVDaOmzejNrPZa9qY6YpKtNM7a4N1SzUiMT5d34Hnc");
  console.log(localStorage.getItem('jwtToken'));
  const payload: JwtPayload = jwtDecode<JwtPayload>(localStorage.getItem('jwtToken') as any);
  console.log(payload.id);
  return (
    <>
      <div className={cl['container-modalWindow']}>
        <div className={cl['container-modalWindow_nickname']}>
          Smallgigach
        </div>
        <div className={cl['container-modalWindow_links']}>
          <NavLink to={`/user?id=${payload.id}`} className={cl['container-modalWindow_links__title']}>Мой профиль</NavLink>
        </div>
        <div className={cl['container-modalWindow_links']}>
          <NavLink to={`/myfavourite`} className={cl['container-modalWindow_links__title']}>Мои книги</NavLink>
        </div>
        <div className={cl['container-modalWindow_links']}>
          <NavLink to={`/`} className={cl['container-modalWindow_links__title']}>Настройки</NavLink>
        </div>
        <div className={cl['container-modalWindow_links']}>
          <NavLink to={`/chat`} className={cl['container-modalWindow_links__title']}>Чат с поддержкой</NavLink>
        </div>
        <div className={cl['container-modalWindow_exit']}>
          <NavLink to={`/chat`} className={cl['container-modalWindow_links__title__exit']}>Выход</NavLink>
        </div>
      </div>
    </>
  );
}