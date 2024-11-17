import { MainPage } from "./pages/MainPage/MainPage.tsx";
import { Route, Routes } from "react-router-dom";
import { SignInPage } from "./pages/signInUpPage/SignInPage.tsx";
import { SignUpPage } from "./pages/signInUpPage/SignUpPage.tsx";
import { BookPage } from "./pages/BookPage/BookPage.tsx";
import { AboutUs } from "./pages/AboutUsPage/AboutUs.tsx";
import Profile from "./pages/ProfilePage/Profile.tsx";
import MyFavourite from "./pages/ProfilePage/pagesComponents/MyFavourite/MyFavourite.tsx";
import MyFriends from "./pages/ProfilePage/pagesComponents/Friends/MyFriends.tsx";
export default function App() {
  return (
    <Routes>
      <Route path="/auth" index element={<SignInPage />} />
      <Route path="/reg" element={<SignUpPage />} />
      <Route path="/" element={<MainPage />} />
      <Route path="/book" element={<BookPage />} />
      <Route path="/aboutUs" element={<AboutUs />}></Route>
      <Route path="/profile" element={<Profile />}></Route>
      <Route  path='/myfavourite' element={<MyFavourite/>}></Route>
      <Route path="/myfriends" element={<MyFriends/>}></Route>
    </Routes>
  );
}
