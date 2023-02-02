import '../../styles/glow.css'
import Input from './Input'
import ScaleIcon from './ScaleIcon'
import { useRef } from 'react'

const CreateImageBar = () => {
  const widthRef = useRef<HTMLInputElement>(null)
  const heightRef = useRef<HTMLInputElement>(null)
  const titleRef = useRef<HTMLInputElement>(null)

  const handleDownload = async () => {
    if (!widthRef.current || !heightRef.current || !titleRef.current) return
    const width = widthRef.current.value
    const height = heightRef.current.value
    const title = titleRef.current.value
      .replace(/\s{2,}/g, ' ')
      .split(' ')
      .join('-')

    const response = await fetch(
      `https://placeholder-image-production.up.railway.app/${width}x${height}?text=${title}`
    )

    const imageBlob = await response.blob()

    const imageURL = URL.createObjectURL(imageBlob)

    const link = document.createElement('a')
    link.style.display = 'none'
    link.href = imageURL
    link.download = `${title}.png`
    document.body.appendChild(link)
    link.click()
    window.URL.revokeObjectURL(imageURL)
    document.body.removeChild(link)
  }
  return (
    <>
      <div
        className={
          'py-5 px-5 sm:px-12 flex flex-col sm:flex-row items-center justify-center bg-[#201D24] text-[#B9B4C7] border border-[#494253] mt-32 rounded-lg sm:rounded-full max-w-xl glow-tab font-medium gap-3'
        }
      >
        <div className={'flex gap-4 items-center'}>
          <Input
            ref={widthRef}
            type={'number'}
            className={'text-center w-[80px]'}
          />

          <ScaleIcon />

          <Input
            ref={heightRef}
            type={'number'}
            className={'text-center w-[80px]'}
          />
        </div>

        <Input
          ref={titleRef}
          type={'text'}
          className={'w-full sm:w-[200px]'}
          placeholder={'my amazing photo'}
        />
      </div>
      <button
        onClick={handleDownload}
        className={
          'mt-12 bg-[#8C3DF1] rounded-md font-bold text-white px-6 py-3 hover:bg-[#7330C9]'
        }
      >
        Download
      </button>
    </>
  )
}

export default CreateImageBar
