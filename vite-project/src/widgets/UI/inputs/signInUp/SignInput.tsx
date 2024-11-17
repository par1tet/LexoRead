import cl from "./SignInput.module.css";

type SignInput = {
    type: string,
    placeholder: string
}

export default function SignInput({type, placeholder}: SignInput) {
  return (
    <>
        <input type={type} placeholder={placeholder} className={cl['input']}  onChange={(e) => console.log(e.target.value)} />
    </>
  );
}
