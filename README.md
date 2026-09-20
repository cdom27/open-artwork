# Open Artwork

Open Artwork is an open art data project for exploring artists, artworks, metadata,
images, and color palettes. It is primarily a portfolio and learning project for
exploring frontend development, backend architecture, and image-processing work.

The rewrite is currently in an early stage: the architecture, data model, and
public API are subject to change.

## Project status

The existing implementation is a legacy Express and TypeScript API. It remains
useful as a reference for the current data and endpoints, but it is not the
long-term architecture for the project.

The rewrite will become the primary version of Open Artwork once its core
workflows and API contract are stable. The project is not currently deployed;
the API will return when the new architecture is complete.

## Rewrite direction

### Frontend

The frontend will be built with Svelte. Open Artwork is content-heavy, so the rewrite prioritizes:

- Server-side rendering and pre-rendering
- Artwork metadata and artist pages
- Image galleries and artwork palettes
- Search and discovery
- Clear API documentation

### Backend

The backend will be built with Go and the Gin web framework. Go gives the project a useful place outside of simple CRUD operations to explore CPU-oriented work, including:

- Image processing
- Palette generation
- Data ingestion and transformation

The Go backend will eventually replace the legacy Express application.

## Goals

The rewrite is intended to:

1. Make Open Artwork a better way to browse and understand public-domain art.
2. Provide a reliable, well-documented API for developers and creative projects.
3. Treat images, metadata, and palettes as first-class parts of the product.
4. Create a focused environment for exploring image-processing workloads.
5. Keep the codebase approachable for contributors and useful as a portfolio
   project.

## Legacy API

The legacy Express API is currently offline and has been for some time. It is kept in the repository as a reference for the previous data model and endpoint design, not as an active service.

The API is not available for requests, and no API key can currently be generated. The rewrite may introduce a different base URL, authentication model, data model, and endpoint structure.

The previous API supported artist metadata, artwork metadata, and public-domain image URLs. Its old documentation is not a contract for the rewrite. Open Artwork focuses on public-domain artwork and openly available metadata. Sources and attribution will be documented as the rewrite's data model and ingestion pipeline are established.

Do not assume that every field, image URL, or endpoint in the legacy API will remain available in the rewrite.
