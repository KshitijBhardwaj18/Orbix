"use client"
import React, { useState, useEffect } from "react";

const carouselData = [
  {
    id: 1,
    title: "Jai hind",
    subtitle: "our platform was born in bharat",
    image: "/indiagate.jpg"
  },
  {
    id: 2,
    title: "Trade with confidence",
    subtitle: "Advanced trading tools at your fingertips",
    image: "/mountain.jpg"
  },
  {
    id: 3,
    title: "Secure your investments",
    subtitle: "Bank-level security for your assets",
    image: "/earth.jpg"
  },
  {
    id: 4,
    title: "Join the revolution",
    subtitle: "Be part of the crypto future",
    image: "/space.jpg"
  },
  {
    id: 5,
    title: "Start trading today",
    subtitle: "No fees for the first month",
    image: "/mountain.jpg"
  }
];

const Thumbnail = () => {
  const [currentIndex, setCurrentIndex] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentIndex((prevIndex) => 
        prevIndex === carouselData.length - 1 ? 0 : prevIndex + 1
      );
    }, 4000);

    return () => clearInterval(interval);
  }, []);

  const goToSlide = (index: number) => {
    setCurrentIndex(index);
  };

  return (
    <div className="relative mb-5 h-72 w-full rounded-2xl overflow-hidden">
      {/* All slides stacked on top of each other */}
      {carouselData.map((slide, index) => (
        <div 
          key={slide.id} 
          className={`absolute inset-0 w-full h-full transition-transform duration-700 ease-in-out ${
            index === currentIndex ? 'translate-x-0' : 
            index < currentIndex ? '-translate-x-full' : 'translate-x-full'
          }`}
        >
          {index == 0 ? <img
            alt={slide.title}
            src={slide.image}
            className="h-full w-full rounded-2xl object-cover "
          /> : <img
          alt={slide.title}
          src={slide.image}
          className="h-full w-full rounded-2xl object-cover object-bottom"
        /> }
          <div
            className="absolute inset-0 rounded-2xl"
            style={{
              background: `
                radial-gradient(circle at center, 
                    transparent 40%, 
                    transparent 10%, 
                    rgba(0,0,0,0.9) 70%, 
                    rgba(0,0,0,0.8) 100%
                )
              `,
            }}
          />
          
          {/* Content for each slide */}
          <div className="absolute bottom-5 left-5 z-10">
            <p className="text-2xl font-bold text-white font-serif">
              {slide.title}
            </p>
            {slide.subtitle && (
              <p className="text-sm text-gray-300 mt-2">
                {slide.subtitle}
              </p>
            )}
          </div>
        </div>
      ))}

      {/* Dots indicator */}
      <div className="absolute bottom-5 right-5 z-20 flex space-x-2">
        {carouselData.map((_, index) => (
          <button
            key={index}
            onClick={() => goToSlide(index)}
            className={`w-2 h-2 rounded-full transition-all duration-300 ${
              index === currentIndex 
                ? 'bg-white' 
                : 'bg-white/40 hover:bg-white/60'
            }`}
          />
        ))}
      </div>

      {/* Progress bar
      <div className="absolute top-0 left-0 w-full h-1 bg-black/20 z-20">
        <div 
          className="h-full bg-white/60 transition-all duration-100 ease-linear"
          style={{ 
            width: `${((currentIndex + 1) / carouselData.length) * 100}%` 
          }}
        />
      </div> */}

      {/* Navigation arrows */}
      <button 
        onClick={() => goToSlide(currentIndex === 0 ? carouselData.length - 1 : currentIndex - 1)}
        className="absolute left-3 top-1/2 transform -translate-y-1/2 z-20 bg-black/50 hover:bg-black/70 text-white p-2 rounded-full transition-all duration-200"
      >
        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
        </svg>
      </button>
      
      <button 
        onClick={() => goToSlide(currentIndex === carouselData.length - 1 ? 0 : currentIndex + 1)}
        className="absolute right-3 top-1/2 transform -translate-y-1/2 z-20 bg-black/50 hover:bg-black/70 text-white p-2 rounded-full transition-all duration-200"
      >
        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
        </svg>
      </button>
    </div>
  );
};

export default Thumbnail;