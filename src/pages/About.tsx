import Layout from "@/components/layout/Layout";
import { Users, Award, Truck, Clock } from "lucide-react";

const stats = [
  { icon: Users, value: "50,000+", label: "Happy Customers" },
  { icon: Award, value: "10+", label: "Years Experience" },
  { icon: Truck, value: "100+", label: "Cities Served" },
  { icon: Clock, value: "4 Hrs", label: "Express Delivery" },
];

const About = () => {
  return (
    <Layout>
      <div className="gradient-primary py-16 text-primary-foreground">
        <div className="container text-center">
          <h1 className="text-4xl md:text-5xl font-heading font-bold mb-4">About LK Printers</h1>
          <p className="text-xl opacity-90 max-w-2xl mx-auto">
            Your trusted partner for all printing needs since 2014
          </p>
        </div>
      </div>

      <div className="container py-12">
        <div className="grid md:grid-cols-4 gap-6 mb-16">
          {stats.map((stat, index) => {
            const Icon = stat.icon;
            return (
              <div key={index} className="bg-card rounded-xl border border-border p-6 text-center">
                <Icon className="w-10 h-10 mx-auto text-primary mb-3" />
                <div className="text-3xl font-bold text-foreground mb-1">{stat.value}</div>
                <div className="text-muted-foreground text-sm">{stat.label}</div>
              </div>
            );
          })}
        </div>

        <div className="max-w-3xl mx-auto space-y-8">
          <section>
            <h2 className="text-2xl font-heading font-bold mb-4">Our Story</h2>
            <p className="text-muted-foreground leading-relaxed">
              LK Printers started in 2014 with a simple mission: to make professional printing accessible 
              to everyone. What began as a small printing shop has grown into one of India's leading 
              online printing platforms, serving over 50,000 customers across 100+ cities.
            </p>
          </section>

          <section>
            <h2 className="text-2xl font-heading font-bold mb-4">Our Mission</h2>
            <p className="text-muted-foreground leading-relaxed">
              We believe that every business, big or small, deserves access to high-quality printing 
              at affordable prices. Our mission is to simplify the printing process and deliver 
              exceptional products that help our customers make a lasting impression.
            </p>
          </section>

          <section>
            <h2 className="text-2xl font-heading font-bold mb-4">Why Choose Us?</h2>
            <ul className="space-y-3 text-muted-foreground">
              <li className="flex items-start gap-3">
                <div className="w-2 h-2 rounded-full bg-primary mt-2" />
                <span>Premium quality printing with state-of-the-art equipment</span>
              </li>
              <li className="flex items-start gap-3">
                <div className="w-2 h-2 rounded-full bg-primary mt-2" />
                <span>Express delivery options including same-day delivery</span>
              </li>
              <li className="flex items-start gap-3">
                <div className="w-2 h-2 rounded-full bg-primary mt-2" />
                <span>Easy-to-use online design tools and templates</span>
              </li>
              <li className="flex items-start gap-3">
                <div className="w-2 h-2 rounded-full bg-primary mt-2" />
                <span>Dedicated customer support team</span>
              </li>
              <li className="flex items-start gap-3">
                <div className="w-2 h-2 rounded-full bg-primary mt-2" />
                <span>Competitive pricing with bulk discounts</span>
              </li>
            </ul>
          </section>
        </div>
      </div>
    </Layout>
  );
};

export default About;
