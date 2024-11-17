import avatar from "../../../assets/img/avatar.png";
import cl from "../AvatarAndNicknameForProfilePage/AvatarAndNickname.module.css";
type AvatarAndNicknameProps = {
  nickname: string;
  friends: string;
};
const AvatarAndNick = ({nickname, friends}: AvatarAndNicknameProps) => {
  return (
    <>
      <div className={cl["container-avatar"]}>
        <div className={cl["header-logo"]}>
          <img src={avatar} alt="avatarUrl" />
        </div>
        <div className={cl['container-avatar_friends']}>{friends}</div>
        <div className={cl["username"]}>{nickname}</div>
      </div>
    </>
  );
};

export default AvatarAndNick;
