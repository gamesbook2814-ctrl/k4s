import { Link } from "react-router-dom";
import { Facebook, Twitter, Instagram, Linkedin, Youtube, Phone, Mail, MapPin } from "lucide-react";

const Footer = () => {
  return (
    <footer className="bg-foreground text-background">
      <div className="container py-12">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-8">
          {/* Company Info */}
          <div className="lg:col-span-2">
            <h2 className="text-2xl font-heading font-bold mb-4">
              <span className="text-accent">LK</span> Printers
            </h2>
            <p className="text-background/70 mb-4 text-sm leading-relaxed">
              Your one-stop destination for all printing needs. From business cards to corporate gifts, 
              we deliver quality printing solutions with fast turnaround times.
            </p>
            <div className="flex gap-4">
              <a href="#" className="text-background/70 hover:text-accent transition-colors">
                <Facebook className="w-5 h-5" />
              </a>
              <a href="#" className="text-background/70 hover:text-accent transition-colors">
                <Twitter className="w-5 h-5" />
              </a>
              <a href="#" className="text-background/70 hover:text-accent transition-colors">
                <Instagram className="w-5 h-5" />
              </a>
              <a href="#" className="text-background/70 hover:text-accent transition-colors">
                <Linkedin className="w-5 h-5" />
              </a>
              <a href="#" className="text-background/70 hover:text-accent transition-colors">
                <Youtube className="w-5 h-5" />
              </a>
            </div>
          </div>

          {/* Quick Links */}
          <div>
            <h3 className="font-heading font-semibold mb-4">Quick Links</h3>
            <ul className="space-y-2 text-sm">
              <li><Link to="/about" className="text-background/70 hover:text-accent transition-colors">About Us</Link></li>
              <li><Link to="/products" className="text-background/70 hover:text-accent transition-colors">All Products</Link></li>
              <li><Link to="/same-day-delivery" className="text-background/70 hover:text-accent transition-colors">Same Day Delivery</Link></li>
              <li><Link to="/bulk-buying" className="text-background/70 hover:text-accent transition-colors">Bulk Orders</Link></li>
              <li><Link to="/blog" className="text-background/70 hover:text-accent transition-colors">Blog</Link></li>
            </ul>
          </div>

          {/* Categories */}
          <div>
            <h3 className="font-heading font-semibold mb-4">Categories</h3>
            <ul className="space-y-2 text-sm">
              <li><Link to="/stationery" className="text-background/70 hover:text-accent transition-colors">Business Stationery</Link></li>
              <li><Link to="/apparel" className="text-background/70 hover:text-accent transition-colors">Custom Apparel</Link></li>
              <li><Link to="/packaging" className="text-background/70 hover:text-accent transition-colors">Packaging</Link></li>
              <li><Link to="/corporate-gifts" className="text-background/70 hover:text-accent transition-colors">Corporate Gifts</Link></li>
              <li><Link to="/photo-gifts" className="text-background/70 hover:text-accent transition-colors">Photo Gifts</Link></li>
            </ul>
          </div>

          {/* Contact */}
          <div>
            <h3 className="font-heading font-semibold mb-4">Contact Us</h3>
            <ul className="space-y-3 text-sm">
              <li className="flex items-start gap-2">
                <Phone className="w-4 h-4 mt-0.5 text-accent" />
                <span className="text-background/70">+91 1800-123-4567</span>
              </li>
              <li className="flex items-start gap-2">
                <Mail className="w-4 h-4 mt-0.5 text-accent" />
                <span className="text-background/70">support@lkprinters.com</span>
              </li>
              <li className="flex items-start gap-2">
                <MapPin className="w-4 h-4 mt-0.5 text-accent" />
                <span className="text-background/70">123 Print Street, Business Park, India</span>
              </li>
            </ul>
          </div>
        </div>

        {/* Bottom Bar */}
        <div className="border-t border-background/20 mt-10 pt-6 flex flex-col md:flex-row justify-between items-center gap-4">
          <p className="text-sm text-background/60">
            © 2024 LK Printers. All rights reserved.
          </p>
          <div className="flex gap-6 text-sm">
            <Link to="/privacy" className="text-background/60 hover:text-accent transition-colors">Privacy Policy</Link>
            <Link to="/terms" className="text-background/60 hover:text-accent transition-colors">Terms & Conditions</Link>
            <Link to="/refund" className="text-background/60 hover:text-accent transition-colors">Refund Policy</Link>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
