import Link from 'next/link';
import { ArrowRight, Database, Zap, Code } from 'lucide-react';

export default function HomePage() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-background">
      <div className="max-w-5xl mx-auto px-6 py-20 text-center space-y-12">
        {/* Hero Section */}
        <div className="space-y-6">
          <div className="inline-flex items-center gap-2 px-4 py-2 rounded-full border border-orange-500/20 bg-orange-500/10 text-orange-600 dark:text-orange-500 text-sm font-medium">
            <Database className="w-4 h-4" />
            In-Memory Database
          </div>
          
          <h1 className="text-6xl md:text-7xl font-bold tracking-tight">
            <span className="bg-gradient-to-r from-orange-500 to-orange-600 bg-clip-text text-transparent">
              ECHO DB
            </span>
          </h1>
          
          <p className="text-xl md:text-2xl text-muted-foreground max-w-2xl mx-auto leading-relaxed">
            A Redis-like in-memory database built on top of RESP protocol
          </p>
          
          <p className="text-base text-muted-foreground/80 max-w-xl mx-auto">
            Built as part of learning how to design and build a Redis-like database. 
            Fast, lightweight, and compatible with existing Redis clients.
          </p>
        </div>

        {/* CTA Buttons */}
        <div className="flex flex-wrap gap-4 justify-center">
          <Link
            href="/docs"
            className="inline-flex items-center gap-2 px-6 py-3 bg-orange-500 hover:bg-orange-600 text-white rounded-lg font-medium transition-colors shadow-lg shadow-orange-500/20"
          >
            Get Started
            <ArrowRight className="w-4 h-4" />
          </Link>
          <Link
            href="/docs"
            className="inline-flex items-center gap-2 px-6 py-3 border border-border rounded-lg font-medium hover:bg-accent transition-colors"
          >
            Learn More
          </Link>
        </div>

        {/* Features */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 pt-12 max-w-4xl mx-auto">
          <div className="p-6 border rounded-lg hover:border-orange-500/50 transition-colors">
            <Zap className="w-6 h-6 text-orange-500 mb-3" />
            <h3 className="font-semibold mb-2">Fast</h3>
            <p className="text-sm text-muted-foreground">In-memory operations with minimal latency</p>
          </div>
          <div className="p-6 border rounded-lg hover:border-orange-500/50 transition-colors">
            <Code className="w-6 h-6 text-orange-500 mb-3" />
            <h3 className="font-semibold mb-2">RESP Protocol</h3>
            <p className="text-sm text-muted-foreground">Compatible with Redis clients and tools</p>
          </div>
          <div className="p-6 border rounded-lg hover:border-orange-500/50 transition-colors">
            <Database className="w-6 h-6 text-orange-500 mb-3" />
            <h3 className="font-semibold mb-2">Simple</h3>
            <p className="text-sm text-muted-foreground">Easy to install, configure, and use</p>
          </div>
        </div>
      </div>
    </div>
  );
}
