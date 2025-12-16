import { useState, useEffect } from "react";
import { ChevronLeft, ChevronRight } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";

const slides = [
  {
    id: 1,
    title: "Begin 2026 in Style",
    subtitle: "Select from our latest collection of exclusive Diaries and Calendars",
    cta: "Shop Now",
    href: "/calendars-diaries",
    gradient: "from-primary via-primary-dark to-primary",
  },
  {
    id: 2,
    title: "Custom Apparel",
    subtitle: "Premium quality custom printed T-shirts, Hoodies & more",
    cta: "Explore Collection",
    href: "/apparel",
    gradient: "from-accent via-accent-dark to-accent",
  },
  {
    id: 3,
    title: "Same Day Delivery",
    subtitle: "Get your prints delivered within 4 hours in major cities",
    cta: "Order Now",
    href: "/same-day-delivery",
    gradient: "from-success via-emerald-600 to-success",
  },
  {
    id: 4,
    title: "Corporate Gifts",
    subtitle: "Premium gifting solutions for your business partners",
    cta: "View Gifts",
    href: "/corporate-gifts",
    gradient: "from-primary via-violet-600 to-primary-dark",
  },
];

const HeroCarousel = () => {
  const [currentSlide, setCurrentSlide] = useState(0);

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentSlide((prev) => (prev + 1) % slides.length);
    }, 5000);
    return () => clearInterval(timer);
  }, []);

  const goToSlide = (index: number) => setCurrentSlide(index);
  const prevSlide = () => setCurrentSlide((prev) => (prev - 1 + slides.length) % slides.length);
  const nextSlide = () => setCurrentSlide((prev) => (prev + 1) % slides.length);

  return (
    <section className="relative overflow-hidden">
      <div className="relative h-[300px] sm:h-[400px] md:h-[500px]">
        {slides.map((slide, index) => (
          <div
            key={slide.id}
            className={`absolute inset-0 transition-all duration-700 ease-in-out ${
              index === currentSlide
                ? "opacity-100 translate-x-0"
                : index < currentSlide
                ? "opacity-0 -translate-x-full"
                : "opacity-0 translate-x-full"
            }`}
          >
            <div className={`h-full bg-gradient-to-r ${slide.gradient} flex items-center`}>
              <div className="container">
                <div className="max-w-xl text-primary-foreground animate-fade-in">
                  <h2 className="text-3xl sm:text-4xl md:text-5xl font-heading font-bold mb-4">
                    {slide.title}
                  </h2>
                  <p className="text-lg sm:text-xl mb-6 opacity-90">
                    {slide.subtitle}
                  </p>
                  <Link to={slide.href}>
                    <Button variant="hero" size="xl">
                      {slide.cta}
                    </Button>
                  </Link>
                </div>
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* Navigation Arrows */}
      <button
        onClick={prevSlide}
        className="absolute left-4 top-1/2 -translate-y-1/2 bg-background/20 hover:bg-background/40 text-primary-foreground p-2 rounded-full transition-colors"
      >
        <ChevronLeft className="w-6 h-6" />
      </button>
      <button
        onClick={nextSlide}
        className="absolute right-4 top-1/2 -translate-y-1/2 bg-background/20 hover:bg-background/40 text-primary-foreground p-2 rounded-full transition-colors"
      >
        <ChevronRight className="w-6 h-6" />
      </button>

      {/* Dots */}
      <div className="absolute bottom-4 left-1/2 -translate-x-1/2 flex gap-2">
        {slides.map((_, index) => (
          <button
            key={index}
            onClick={() => goToSlide(index)}
            className={`w-3 h-3 rounded-full transition-all ${
              index === currentSlide
                ? "bg-primary-foreground scale-110"
                : "bg-primary-foreground/50 hover:bg-primary-foreground/70"
            }`}
          />
        ))}
      </div>
    </section>
  );
};

export default HeroCarousel;
