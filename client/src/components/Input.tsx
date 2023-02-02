import { ForwardedRef, forwardRef, RefObject } from 'react'

interface IInputProps {
  className?: string
  type: string
  placeholder?: string
}

const Input = forwardRef(
  (
    { className, type, placeholder }: IInputProps,
    ref: ForwardedRef<HTMLInputElement>
  ) => {
    const classes = `bg-[#302D39] hover:bg-[#2C2A33] py-1 px-4 rounded-md  h-9 focus:outline-0 placeholder-[#63616A] ${className}`

    return (
      <input
        className={classes}
        ref={ref}
        type={type}
        placeholder={placeholder}
        spellCheck={'false'}
      />
    )
  }
)

export default Input
