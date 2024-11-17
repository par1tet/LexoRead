import AvatarAndNick from "../../../../widgets/pageComponents/AvatarAndNicknameForProfilePage/AvatarAndNick";
import {Header} from "../../../../widgets/pageComponents/Header";
import cl from "../Friends/MyFriend.module.css";
import FriendsList from "../../../../widgets/pageComponents/FrendsList/FriendsList.tsx";

const MyFriends = () => {
    type User = {
        coverPath: string
        username: string
    }
    const users: User[] = [{
        username: 'zxc',
        coverPath: 'https://www.interfax.ru/ftproot/photos/photostory/2019/07/09/week4_700.jpg'
    },
        {
            username: 'cxcxcx',
            coverPath: 'https://www.interfax.ru/ftproot/photos/photostory/2019/07/09/week4_700.jpg'
        },
        {
            username: 'cxcxcx',
            coverPath: 'https://www.interfax.ru/ftproot/photos/photostory/2019/07/09/week4_700.jpg'
        },
        {
            username: 'cxcxcx',
            coverPath: 'https://www.interfax.ru/ftproot/photos/photostory/2019/07/09/week4_700.jpg'
        },
        {
            username: 'cxcxcx',
            coverPath: 'https://www.interfax.ru/ftproot/photos/photostory/2019/07/09/week4_700.jpg'
        },
        {
            username: 'cxcxcx',
            coverPath: 'https://www.interfax.ru/ftproot/photos/photostory/2019/07/09/week4_700.jpg'
        },
        {
            username: 'cxcxcx',
            coverPath: 'https://www.interfax.ru/ftproot/photos/photostory/2019/07/09/week4_700.jpg'
        },
        {
            username: 'cxcxcx',
            coverPath: 'https://www.interfax.ru/ftproot/photos/photostory/2019/07/09/week4_700.jpg'
        },
        {
            username: 'cxcxcx',
            coverPath: 'https://www.interfax.ru/ftproot/photos/photostory/2019/07/09/week4_700.jpg'
        },
        {
            username: 'cxcxcx',
            coverPath: 'https://www.interfax.ru/ftproot/photos/photostory/2019/07/09/week4_700.jpg'
        },
        {
            username: 'cxcxcx',
            coverPath: 'https://www.interfax.ru/ftproot/photos/photostory/2019/07/09/week4_700.jpg'
        },
        {
            username: 'cxcxcx',
            coverPath: 'https://www.interfax.ru/ftproot/photos/photostory/2019/07/09/week4_700.jpg'
        },

    ]
    return (
        <>
            <Header></Header>
            <AvatarAndNick
                nickname={"smallgigach"}
                friends={"Друзья"}
            ></AvatarAndNick>
            <div className={cl["container-friendsList"]}>
                <div className={cl['container-friends_']}>

                </div>
                <FriendsList users={users}/>
            </div>
        </>
    );
};

export default MyFriends;
