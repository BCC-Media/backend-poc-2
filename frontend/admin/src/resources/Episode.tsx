import { Datagrid, List, NumberField, TextField, DateField, ListProps, EditButton, SelectInput, Edit, SimpleForm, SaveButton, EditProps, CreateProps, Create, TextInput, NumberInput, ReferenceInput, SelectArrayInput, ReferenceField, FormDataConsumer, SimpleFormView, TabbedForm, Tab, FormTab, FormWithRedirect, DateTimeInput, ReferenceManyField, ArrayField, useWarnWhenUnsavedChanges, Toolbar, createMuiTheme, Link, TopToolbar, CreateButton, ExportButton, Button } from 'react-admin';
// in src/App.js
import React, { cloneElement } from 'react';
import { AgeRatingChoices } from '../models/AgeRating';
import ContentAdd from '@material-ui/icons/Add';
import { useLocation } from 'react-router-dom';
import { Media } from '../models/Media';

const ListActions = () => (
    <TopToolbar>
        <CreateButton/>
        <ExportButton/>
    </TopToolbar>
);  

const Total = (props: any) => <div>{props.total}</div>;

export const EpisodeList: React.FC<ListProps> = props => {
    return (
        <List actions={<ListActions/>} {...props}  >
            <Datagrid rowClick="edit">
                <TextField source="mediaType" />
                <TextField source="title" />
                <TextField source="description" />
                <NumberField source="startTime" />
                <NumberField source="endTime" />
                <DateField source="createdAt" />
                <DateField source="updatedAt" />
                <EditButton/>
            </Datagrid>
        </List>
    );
};

export const EpisodeEdit: React.FC<EditProps> = props => {
    return (
        <Edit {...props}>
            <FormWithRedirect warnWhenUnsavedChanges render={formProps =>
                <form>
                    <div className="p-4 flex flex-col">
                        <div className='text-sm'>#{formProps.record?.id} <span className="capitalize">Episode</span></div>
                        <TextField source="title" variant='h6'/>
                        <div>
                        <div className="bg-gray-100 p-2 rounded mt-2 mb-6 text-sm">
                            <span>Created</span> <DateField source="createdAt" showTime />
                            &nbsp;| <span>Last updated </span> <DateField source="updatedAt" showTime />
                        </div></div>
                        <ReferenceInput required source="primaryGroupID" reference="season" label="Belongs to season">
                            <SelectInput optionText="title" />
                        </ReferenceInput>
                        <NumberInput source="sequenceNumber" />
                        <TextInput source="title" />
                        <TextInput source="description" />
                        <TextInput source="longDescription" />
                        <SelectInput source="agerating" choices={AgeRatingChoices}/>
                        <ReferenceInput source="assetID" reference="assets" label="Asset">
                            <SelectInput optionText="name" />
                        </ReferenceInput>
                        
                        <SaveButton
                        saving={formProps.saving}
                        disabled={formProps.pristine}
                        handleSubmitWithRedirect={formProps.handleSubmitWithRedirect}/>

                        <div className="mt-4">
                            <h4>Suclips</h4>
                            <Link to={{
                                pathname: "/subclip/create",
                                state: { initialValues: { primaryGroupID: formProps.record?.id } }
                            }}>
                                <button type="button">
                                    <ContentAdd />Create subclip
                                </button>
                            </Link>
                            <ReferenceManyField label="Episodes" reference="subclip" target="subclippedMediaID">
                                <ArrayField>
                                    <Datagrid rowClick="edit">
                                        <TextField source="id" />
                                        <TextField source="title" />
                                        <TextField source="description" />
                                        <EditButton/>
                                    </Datagrid>
                                </ArrayField>
                            </ReferenceManyField>
                        </div>
                    </div>
                </form>
            }/>
        </Edit>
    )
};

export const EpisodeCreate: React.FC<CreateProps> = props => {
    const location = useLocation<{initialValues: Media}>();
    return (
    <Create {...props}>
        <FormWithRedirect warnWhenUnsavedChanges initialValues={location.state?.initialValues} render={formProps =>
            <form>
                <div className="p-4 flex flex-col">
                    <ReferenceInput required source="primaryGroupID" reference="season" label="Belongs to season">
                        <SelectInput optionText="title" />
                    </ReferenceInput>
                    <NumberInput source="sequenceNumber" />
                    <TextInput source="title" />
                    <TextInput source="description" />
                    <TextInput source="longDescription" />
                    <SelectInput source="agerating" choices={AgeRatingChoices}/>
                    <ReferenceInput source="assetID" reference="assets" label="Asset">
                        <SelectInput optionText="name" />
                    </ReferenceInput>
                    <DateField source="createdAt" showTime />
                    <DateField source="updatedAt" showTime />
                    <Toolbar>
                        <SaveButton
                        saving={formProps.saving}
                        handleSubmitWithRedirect={formProps.handleSubmitWithRedirect}/>
                    </Toolbar>
                </div>
            </form>
        }/>
    </Create>
)};