![Graphe](https://raw.githubusercontent.com/gabrielaravena32/graphe-app/main/build/banner.png)
<br />

## Development & Building

Live development mode: `wails dev`

Build distributable app: `wails build`

## Roadmap

**Version 1 Checklist**

- [x] Add keyboard shortcuts
- [ ] Format display of texts
  - [x] Update virtualiser
  - [ ] Add book titles
  - [ ] Add top padding
- [ ] Get the modals working
  - [ ] 'Choose passage' modal
  - [ ] 'Choose text' modal
  - [ ] 'Search pane' modal
- [ ] Add theme selector to settings in sidebar
- [ ] Hover effect for words + ui for displaying info
- [ ] Get all the bible texts working properly
  - [ ] Split GNT text into word, pre, post (separate out punctuation)
  - [ ] Add HOT
  - [ ] Add ESV
  - [ ] Add LXX

**Later Versions Checklist**

- [ ] Auto-download database file if it does not exist on startup
- [ ] Sidebar: formatting
- [ ] Sidebar: settings
- [ ] Create menubar with full features

**Performance**

- [ ] Split `GetScriptureSections` request into multiple batches for larger ranges that perform concurrently
- [ ] Increase speed of virtualiser when resizing window
