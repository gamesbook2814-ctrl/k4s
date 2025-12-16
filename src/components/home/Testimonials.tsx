import { Star, Quote } from "lucide-react";

const testimonials = [
  {
    id: 1,
    name: "Rajesh Kumar",
    company: "Tech Startup",
    avatar: "/placeholder.svg",
    rating: 5,
    review: "Excellent quality and super fast delivery! LK Printers has been our go-to for all business cards and marketing materials.",
  },
  {
    id: 2,
    name: "Priya Sharma",
    company: "Event Management",
    avatar: "/placeholder.svg",
    rating: 5,
    review: "The same-day delivery service is a lifesaver for our last-minute event needs. Highly recommend!",
  },
  {
    id: 3,
    name: "Amit Patel",
    company: "Restaurant Owner",
    avatar: "/placeholder.svg",
    rating: 5,
    review: "Beautiful menu designs and packaging solutions. The quality exceeded our expectations every time.",
  },
];

const Testimonials = () => {
  return (
    <section className="py-12 md:py-16">
      <div className="container">
        <h2 className="text-2xl md:text-3xl font-heading font-bold text-center mb-3">
          What Our Customers Say
        </h2>
        <p className="text-muted-foreground text-center mb-10 max-w-2xl mx-auto">
          Join thousands of satisfied customers who trust LK Printers for their printing needs
        </p>

        <div className="grid md:grid-cols-3 gap-6">
          {testimonials.map((testimonial) => (
            <div
              key={testimonial.id}
              className="bg-card rounded-xl p-6 border border-border hover:shadow-lg transition-shadow relative"
            >
              <Quote className="absolute top-4 right-4 w-8 h-8 text-primary/10" />
              <div className="flex items-center gap-3 mb-4">
                <img
                  src={testimonial.avatar}
                  alt={testimonial.name}
                  className="w-12 h-12 rounded-full object-cover bg-muted"
                />
                <div>
                  <h4 className="font-semibold text-foreground">{testimonial.name}</h4>
                  <p className="text-sm text-muted-foreground">{testimonial.company}</p>
                </div>
              </div>
              <div className="flex gap-1 mb-3">
                {Array.from({ length: testimonial.rating }).map((_, i) => (
                  <Star key={i} className="w-4 h-4 fill-amber-400 text-amber-400" />
                ))}
              </div>
              <p className="text-muted-foreground text-sm leading-relaxed">
                "{testimonial.review}"
              </p>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};

export default Testimonials;
