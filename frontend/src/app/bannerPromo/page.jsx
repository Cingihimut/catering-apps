"use client"

import { ArrowCircleLeft, ArrowCircleRight } from "@phosphor-icons/react"
import { useState } from "react"

const Carousel = ({ children: slides }) => {
    const [current, setCurrent] = useState()

    const prev = () =>
        setCurrent((current) => (current === 0 ? slides.length -1 : current -1
            ))
    const next = () =>
        setCurrent((current) => (current === slides.length -1 ? 0 : current +1
            ))

    return(
        <div className="overflow-hiden relative">
            <div className="flex">{slides}</div>
            <div className="absolute inset-0 items-center justify-between p-4">
                <button onClick={prev} className="p-1 rounded-full shadow bg-white-800 text-gray-800 hover:bg-white">
                    <ArrowCircleLeft size={32} />
                </button>
                <button onClick={next} className="p-1 rounded-full shadow bg-white-800 text-gray-800 hover:bg-white">
                    <ArrowCircleRight size={32} />
                </button>
            </div>
        </div>
    )
}
export default Carousel