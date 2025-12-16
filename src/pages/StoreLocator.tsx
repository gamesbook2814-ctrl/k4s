import Layout from "@/components/layout/Layout";
import { MapPin, Phone, Clock } from "lucide-react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useState } from "react";

const stores = [
  {
    id: 1,
    name: "LK Printers - Bangalore HQ",
    address: "123 MG Road, Bangalore, Karnataka 560001",
    phone: "+91 80-1234-5678",
    hours: "Mon-Sat: 9 AM - 8 PM",
    city: "Bangalore",
  },
  {
    id: 2,
    name: "LK Printers - Hyderabad",
    address: "456 Hitech City, Hyderabad, Telangana 500081",
    phone: "+91 40-9876-5432",
    hours: "Mon-Sat: 9 AM - 8 PM",
    city: "Hyderabad",
  },
  {
    id: 3,
    name: "LK Printers - Chennai",
    address: "789 Anna Nagar, Chennai, Tamil Nadu 600040",
    phone: "+91 44-5678-1234",
    hours: "Mon-Sat: 9 AM - 8 PM",
    city: "Chennai",
  },
  {
    id: 4,
    name: "LK Printers - Delhi",
    address: "321 Connaught Place, New Delhi 110001",
    phone: "+91 11-4321-8765",
    hours: "Mon-Sat: 9 AM - 8 PM",
    city: "Delhi",
  },
  {
    id: 5,
    name: "LK Printers - Mumbai",
    address: "555 Andheri West, Mumbai, Maharashtra 400053",
    phone: "+91 22-6789-0123",
    hours: "Mon-Sat: 9 AM - 8 PM",
    city: "Mumbai",
  },
];

const StoreLocator = () => {
  const [searchCity, setSearchCity] = useState("");

  const filteredStores = stores.filter(store =>
    store.city.toLowerCase().includes(searchCity.toLowerCase()) ||
    store.address.toLowerCase().includes(searchCity.toLowerCase())
  );

  return (
    <Layout>
      <div className="container py-8">
        <h1 className="text-3xl font-heading font-bold mb-8">Find a Store</h1>

        <div className="flex gap-4 mb-8 max-w-md">
          <Input
            placeholder="Search by city or location..."
            value={searchCity}
            onChange={(e) => setSearchCity(e.target.value)}
          />
          <Button variant="accent">Search</Button>
        </div>

        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
          {filteredStores.map((store) => (
            <div key={store.id} className="bg-card rounded-xl border border-border p-6">
              <h3 className="font-heading font-semibold text-lg mb-3">{store.name}</h3>
              <div className="space-y-2 text-sm">
                <div className="flex items-start gap-2 text-muted-foreground">
                  <MapPin className="w-4 h-4 mt-0.5 flex-shrink-0 text-primary" />
                  <span>{store.address}</span>
                </div>
                <div className="flex items-center gap-2 text-muted-foreground">
                  <Phone className="w-4 h-4 flex-shrink-0 text-primary" />
                  <a href={`tel:${store.phone}`} className="hover:text-primary">{store.phone}</a>
                </div>
                <div className="flex items-center gap-2 text-muted-foreground">
                  <Clock className="w-4 h-4 flex-shrink-0 text-primary" />
                  <span>{store.hours}</span>
                </div>
              </div>
              <Button variant="outline" size="sm" className="mt-4 w-full">
                Get Directions
              </Button>
            </div>
          ))}
        </div>

        {filteredStores.length === 0 && (
          <p className="text-center text-muted-foreground py-12">
            No stores found in your area. Try a different search.
          </p>
        )}
      </div>
    </Layout>
  );
};

export default StoreLocator;
