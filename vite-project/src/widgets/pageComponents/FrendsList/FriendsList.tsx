import cl from "../FrendsList/FriendsList.module.css";
import {useNavigate} from "react-router-dom";

type listFriendsProps = {
    users: {
        username: string;
        coverPath: string;
    }[];
};
const FriendsList = ({users}: listFriendsProps) => {
    const navigate = useNavigate()
    return (
        <>
            <div className={cl["container_ListFriends"]}>
                <div className={cl['ListFriends_friends']} >
                    {users.map((user, index: number) =>
                        <div key={index} className={cl['friends']}>
                            <img src={user.coverPath} key={index} className={cl['avatarImg']}
                                 onClick={() => {
                                     navigate("/friend?id=2");
                                 }}/>
                            <span className={cl['nickname']}>{user.username}</span>
                        </div>
                    )}
                    
                </div>
            </div>
        </>
    );
};

export default FriendsList;
