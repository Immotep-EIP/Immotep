import React, { useEffect, useState } from 'react'
import { useTranslation } from 'react-i18next'
import i18n from 'i18next'
import {
  Button,
  Image,
  Segmented,
  Upload,
  GetProp,
  UploadFile,
  UploadProps,
  Input,
  message
} from 'antd'
import { PlusOutlined, LogoutOutlined } from '@ant-design/icons'

import { useAuth } from '@/context/authContext'
import SubtitledElement from '@/components/SubtitledElement/SubtitledElement'
import EditIcon from '@/assets/icons/edit.png'
import SaveIcon from '@/assets/icons/save.png'
import CloseIcon from '@/assets/icons/close.png'
import UpdateUserInfos from '@/services/api/User/UpdateUserInfos'
import style from './Settings.module.css'

type FileType = Parameters<GetProp<UploadProps, 'beforeUpload'>>[0]

const getBase64 = (file: FileType): Promise<string> =>
  new Promise((resolve, reject) => {
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = () => resolve(reader.result as string)
    reader.onerror = error => reject(error)
  })

interface UserSettingsProps {
  t: (key: string) => string
}

const UserSettings: React.FC<UserSettingsProps> = ({ t }) => {
  const { user, updateUser } = useAuth()

  const [previewOpen, setPreviewOpen] = useState(false)
  const [previewImage, setPreviewImage] = useState('')
  const [fileList, setFileList] = useState<UploadFile[]>([
    {
      uid: '-1',
      name: 'image.png',
      status: 'done',
      url: 'https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png'
    }
  ])
  const [editData, setEditData] = useState(false)
  const [oldData, setOldData] = useState({
    firstname: user?.firstname,
    lastname: user?.lastname,
    email: user?.email
  })
  const [newData, setNewData] = useState({
    firstname: user?.firstname,
    lastname: user?.lastname,
    email: user?.email
  })

  const saveNewData = async () => {
    if (
      newData.firstname === oldData.firstname &&
      newData.lastname === oldData.lastname
    ) {
      setEditData(false);
      message.info(t('components.messages.no_modifications'));
      return;
    }
    try {
      await UpdateUserInfos({
        firstname: newData.firstname as string,
        lastname: newData.lastname as string,
      });
      setOldData(newData);
      updateUser(newData);
      setEditData(false);
      message.success(t('components.messages.modifications_saved'));
    } catch (error) {
      console.error('Error updating user data:', error);
    }
  };

  const cancelEdit = () => {
    setNewData(oldData);
    setEditData(false);
  };

  useEffect(() => {
    if (user) {
      setNewData({
        firstname: user.firstname,
        lastname: user.lastname,
        email: user.email,
      });
      setOldData({
        firstname: user.firstname,
        lastname: user.lastname,
        email: user.email,
      });
    }
  }, [user]);

  const handlePreview = async (file: UploadFile) => {
    if (!file.url && !file.preview) {
      const newFile = {
        ...file,
        preview: await getBase64(file.originFileObj as FileType)
      }
      setFileList(prevList =>
        prevList.map(f => (f.uid === newFile.uid ? newFile : f))
      )
    }

    setPreviewImage(file.url || (file.preview as string))
    setPreviewOpen(true)
  }

  const handleChange: UploadProps['onChange'] = ({ fileList: newFileList }) =>
    setFileList(newFileList)

  const uploadButton = (
    <button style={{ border: 0, background: 'none' }} type="button">
      <PlusOutlined />
      <div style={{ marginTop: 8 }}>Upload</div>
    </button>
  )

  return (
    <div className={style.settingsContainer}>
      <div className={style.userItem}>
        <Upload
          action="https://660d2bd96ddfa2943b33731c.mockapi.io/api/upload"
          listType="picture-circle"
          fileList={fileList}
          onPreview={handlePreview}
          onChange={handleChange}
        >
          {fileList.length >= 1 ? null : uploadButton}
        </Upload>
        {previewImage && (
          <Image
            wrapperStyle={{ display: 'none' }}
            preview={{
              visible: previewOpen,
              onVisibleChange: visible => setPreviewOpen(visible),
              afterOpenChange: visible => !visible && setPreviewImage('')
            }}
            src={previewImage}
          />
        )}
      </div>
      <div className={style.userInformations}>
        <div className={style.titleContainer}>
          <b>{t('pages.settings.user_infos')}</b>
          <div className={style.editButtons}>
            {editData && (
              <Button
                type="link"
                style={{ width: 25, height: 25, padding: 10 }}
                onClick={cancelEdit}
              >
                <img
                  src={CloseIcon}
                  alt="edit"
                  style={{ width: 20, height: 20 }}
                />
              </Button>
            )}
            <Button
              type="link"
              style={{ width: 25, height: 25, padding: 10 }}
              onClick={() => editData ? saveNewData() : setEditData(!editData)}
            >
              <img
                src={editData ? SaveIcon : EditIcon}
                alt="edit"
                style={{ width: 20, height: 20 }}
              />
            </Button>
          </div>
        </div>
        <SubtitledElement subtitleKey={t('components.input.first_name.label')} subTitleStyle={{ opacity: 0.6 }}>
          {!editData ? (
            user?.firstname
          ) : (
            <Input
              defaultValue={user?.firstname}
              onChange={e =>
                setNewData({ ...newData, firstname: e.target.value })
              }
            />
          )}
        </SubtitledElement>
        <SubtitledElement subtitleKey={t('components.input.last_name.label')} subTitleStyle={{ opacity: 0.6 }}>
          {!editData ? (
            user?.lastname
          ) : (
            <Input
              defaultValue={user?.lastname}
              onChange={e =>
                setNewData({ ...newData, lastname: e.target.value })
              }
            />
          )}
        </SubtitledElement>
        <SubtitledElement subtitleKey={t('components.input.email.label')} subTitleStyle={{ opacity: 0.6 }}>
          {user?.email}
        </SubtitledElement>
      </div>
    </div>
  )
}

const Settings: React.FC = () => {
  const { t } = useTranslation()
  const { logout } = useAuth()

  const switchLanguage = (language: string) => {
    let lang = ''
    switch (language) {
      case 'fr' as string:
        lang = 'fr'
        localStorage.setItem('lang', 'fr')
        break
      case 'en' as string:
        lang = 'en'
        localStorage.setItem('lang', 'en')
        break
      default:
        lang = 'fr'
        localStorage.setItem('lang', 'fr')
        break
    }
    i18n.changeLanguage(lang)
  }

  return (
    <div className={style.layoutContainer}>
      <UserSettings t={t} />

      <div className={style.settingsContainer}>
        <div className={style.settingsItem}>
          {t('pages.settings.language')}
          <Segmented
            options={[
              { label: t('pages.settings.fr'), value: 'fr' },
              { label: t('pages.settings.en'), value: 'en' }
            ]}
            value={i18n.language}
            onChange={value => switchLanguage(value as string)}
          />
        </div>

        <div className={style.settingsItem}>
          {t('pages.settings.logout')}
          <Button
            type="primary"
            danger
            shape="circle"
            icon={<LogoutOutlined />}
            onClick={() => logout()}
          />
        </div>
      </div>
    </div>
  )
}

export default Settings
