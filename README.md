# Ranku

**Ranku** is a fun side project that lets users vote for their favorite anime characters and see how they stack up on a live leaderboard.

The project was built primarily as a learning exercise: to deepen my understanding of the Go programming language and to revisit Svelte after last using it back in version 3.

## Features
1. **Automated data fetching**: A scheduled cron job that integrates with [AniList’s GraphQL API](https://docs.anilist.co/) to fetch top anime and character data.
2. **Authentication system**: Secure username and password login/signup powered by [Better Auth](https://www.better-auth.com/) using its JWT plugin.
3. **Go-based REST API**: Backend API written in Go, built with the lightweight and idiomatic [chi](https://github.com/go-chi/chi) router.
4. **Real-time leaderboard**: Live updates using Server-Sent Events (SSE), with Redis used for caching and performance.
5. **Modern, responsive UI**: Fully responsive frontend styled with [Tailwind CSS](https://tailwindcss.com/) and componentized using [shadcn-svelte](https://www.shadcn-svelte.com/).



## Screenshots
<img width="1380" height="1013" alt="image" src="https://github.com/user-attachments/assets/ad222183-d7a7-4be5-af6f-ec8cbdd7e1c5" />
<img width="1583" height="1156" alt="image" src="https://github.com/user-attachments/assets/1e15537e-6a62-4e83-ba34-c10ae1029c81" />
<img width="1412" height="1255" alt="image" src="https://github.com/user-attachments/assets/d7a4cdab-c9d8-42f6-8d44-bdb0bfde55b1" />

