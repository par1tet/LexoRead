import { MainPage } from "./pages/MainPage/MainPage.tsx";
import { Route, Routes } from "react-router-dom";
import { SignIn } from "./pages/signInUpPage/SignIn.tsx";
import { SignUp } from "./pages/signInUpPage/SignUp.tsx";
import { AboutUs } from "./pages/AboutUsPage/AboutUs.tsx";
import Profile from "./pages/ProfilePage/Profile.tsx";
import MyFavourite from "./pages/ProfilePage/pagesComponents/MyFavourite.tsx";

export default function App() {
  return (
    <Routes>
      <Route path="/auth" index element={<SignIn />} />
      <Route path="/reg" element={<SignUp />} />
      <Route path="/" element={<MainPage />} />
      <Route path="/aboutUs" element={<AboutUs />}></Route>
      <Route path="/profile" element={<Profile />}></Route>
      <Route path="/myfavourite" element={<MyFavourite />}></Route>
    </Routes>
  );
}
