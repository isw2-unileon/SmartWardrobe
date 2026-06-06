# Design Decisions

## Why Next.js?

- Modern React framework
- Routing support
- Good TypeScript integration
- Server-side rendering capabilities

## Why Go?

- High performance
- Simple concurrency model
- Easy API development
- Efficient memory usage

## Why PostgreSQL?

- Reliable relational database
- Strong consistency
- ACID compliance
- Good Supabase integration

## Why Supabase?

Supabase was selected because it provides:

- Managed PostgreSQL hosting
- Authentication services
- Easy integration with Next.js
- Reduced infrastructure maintenance

## Why CLIP?

CLIP allows automatic classification of:

- garment type
- color
- style

without requiring a custom trained model.

## Why Remove.bg?

Provides accurate background removal while reducing development complexity.

## Service Architecture

The backend follows the Repository Pattern.

Benefits:

- Separation of concerns
- Easier testing
- Better maintainability
- Improved scalability