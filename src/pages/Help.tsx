import Layout from "@/components/layout/Layout";
import { Link } from "react-router-dom";
import { Phone, Mail, MessageCircle, Clock, ChevronRight } from "lucide-react";
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion";

const faqs = [
  {
    question: "What is the delivery time for orders?",
    answer: "Standard delivery takes 3-5 business days. We also offer same-day delivery in major cities for select products.",
  },
  {
    question: "How can I track my order?",
    answer: "You can track your order using the Track Order option in the top menu. Enter your order ID to get real-time updates.",
  },
  {
    question: "What is your return policy?",
    answer: "We offer a 7-day return policy for defective products. Custom printed items cannot be returned unless there's a printing error.",
  },
  {
    question: "Do you offer bulk discounts?",
    answer: "Yes! We offer special pricing for bulk orders. Contact our sales team for custom quotes on large orders.",
  },
  {
    question: "What file formats do you accept?",
    answer: "We accept PDF, AI, PSD, PNG, and JPG files. For best results, we recommend high-resolution PDF files.",
  },
];

const Help = () => {
  return (
    <Layout>
      <div className="container py-8">
        <h1 className="text-3xl font-heading font-bold mb-8">Help Center</h1>

        <div className="grid md:grid-cols-3 gap-6 mb-12">
          <div className="bg-card rounded-xl border border-border p-6 text-center">
            <Phone className="w-10 h-10 mx-auto text-primary mb-4" />
            <h3 className="font-heading font-semibold mb-2">Call Us</h3>
            <p className="text-muted-foreground text-sm mb-4">Mon-Sat, 9 AM - 8 PM</p>
            <a href="tel:18001234567" className="text-primary font-medium">1800-123-4567</a>
          </div>
          <div className="bg-card rounded-xl border border-border p-6 text-center">
            <Mail className="w-10 h-10 mx-auto text-primary mb-4" />
            <h3 className="font-heading font-semibold mb-2">Email Us</h3>
            <p className="text-muted-foreground text-sm mb-4">Get a response within 24 hours</p>
            <a href="mailto:support@lkprinters.com" className="text-primary font-medium">support@lkprinters.com</a>
          </div>
          <div className="bg-card rounded-xl border border-border p-6 text-center">
            <MessageCircle className="w-10 h-10 mx-auto text-primary mb-4" />
            <h3 className="font-heading font-semibold mb-2">Live Chat</h3>
            <p className="text-muted-foreground text-sm mb-4">Chat with our support team</p>
            <button className="text-primary font-medium">Start Chat</button>
          </div>
        </div>

        <div className="max-w-3xl mx-auto">
          <h2 className="text-2xl font-heading font-bold mb-6">Frequently Asked Questions</h2>
          <Accordion type="single" collapsible className="w-full">
            {faqs.map((faq, index) => (
              <AccordionItem key={index} value={`item-${index}`}>
                <AccordionTrigger className="text-left">{faq.question}</AccordionTrigger>
                <AccordionContent className="text-muted-foreground">
                  {faq.answer}
                </AccordionContent>
              </AccordionItem>
            ))}
          </Accordion>
        </div>
      </div>
    </Layout>
  );
};

export default Help;
