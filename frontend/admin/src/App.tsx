// in src/App.js
import * as React from "react";
import { Admin, EditGuesser, ListGuesser, Resource } from 'react-admin';
import { createTheme } from '@mui/material/styles';
import jsonServerProvider from 'ra-data-json-server';
import { MediaList, MediaEdit, MediaCreate } from "./resources/Medias";
import './index.css';
import { Route } from "react-router-dom";
import { MediaType } from "./models/MediaType";
import { Media } from "./models/Media";
import { Menu } from "./components/Menu";
import * as Icons from '@mui/icons-material';
import { AssetEdit, AssetList } from "./resources/Assets";
import { ShowCreate, ShowEdit, ShowList } from "./resources/Show";
import { SeasonCreate, SeasonEdit, SeasonList } from "./resources/Season";
import { EpisodeCreate, EpisodeEdit, EpisodeList } from "./resources/Episode";
import { StandaloneCreate, StandaloneEdit, StandaloneList } from "./resources/Standalone";

const dataProvider = jsonServerProvider('http://localhost:8080');
const App = () => (
    <Admin 
    dataProvider={dataProvider}
    /* customRoutes={[
        <Route exact path="/show" render={MediaListWithFilter({mediaType: MediaType.Show})} />,
        <Route exact path="/season" render={MediaListWithFilter({mediaType: "season"})} />,
        <Route exact path="/episode" render={MediaListWithFilter({mediaType: "episode"})} />,
        <Route exact path="/standalones" render={MediaListWithFilter({mediaType: "standalone"})} />,
    ]} */
    menu={Menu}>
        <Resource name="show" list={ShowList} edit={ShowEdit} create={ShowCreate} />
        <Resource name="season" list={SeasonList} edit={SeasonEdit} create={SeasonCreate} />
        <Resource name="episode" list={EpisodeList} edit={EpisodeEdit} create={EpisodeCreate} />
        <Resource name="standalone" list={StandaloneList} edit={StandaloneEdit} create={StandaloneCreate} />
        <Resource name="subclip" list={MediaList} edit={MediaEdit} create={MediaCreate} />
        <Resource name="assets" list={AssetList} />
        <Resource name="media" edit={MediaEdit} create={MediaCreate} icon={Icons.PermMedia}/>
        <Resource name="asset-versions"/>
    </Admin>
);

export default App;