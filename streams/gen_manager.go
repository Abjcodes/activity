package streams

import (
	propertyaccuracy "github.com/go-fed/activity/streams/impl/activitystreams/property_accuracy"
	propertyactor "github.com/go-fed/activity/streams/impl/activitystreams/property_actor"
	propertyaltitude "github.com/go-fed/activity/streams/impl/activitystreams/property_altitude"
	propertyanyof "github.com/go-fed/activity/streams/impl/activitystreams/property_anyof"
	propertyattachment "github.com/go-fed/activity/streams/impl/activitystreams/property_attachment"
	propertyattributedto "github.com/go-fed/activity/streams/impl/activitystreams/property_attributedto"
	propertyaudience "github.com/go-fed/activity/streams/impl/activitystreams/property_audience"
	propertybcc "github.com/go-fed/activity/streams/impl/activitystreams/property_bcc"
	propertybto "github.com/go-fed/activity/streams/impl/activitystreams/property_bto"
	propertycc "github.com/go-fed/activity/streams/impl/activitystreams/property_cc"
	propertyclosed "github.com/go-fed/activity/streams/impl/activitystreams/property_closed"
	propertycontent "github.com/go-fed/activity/streams/impl/activitystreams/property_content"
	propertycontext "github.com/go-fed/activity/streams/impl/activitystreams/property_context"
	propertycurrent "github.com/go-fed/activity/streams/impl/activitystreams/property_current"
	propertydeleted "github.com/go-fed/activity/streams/impl/activitystreams/property_deleted"
	propertydescribes "github.com/go-fed/activity/streams/impl/activitystreams/property_describes"
	propertyduration "github.com/go-fed/activity/streams/impl/activitystreams/property_duration"
	propertyendtime "github.com/go-fed/activity/streams/impl/activitystreams/property_endtime"
	propertyfirst "github.com/go-fed/activity/streams/impl/activitystreams/property_first"
	propertyfollowers "github.com/go-fed/activity/streams/impl/activitystreams/property_followers"
	propertyfollowing "github.com/go-fed/activity/streams/impl/activitystreams/property_following"
	propertyformertype "github.com/go-fed/activity/streams/impl/activitystreams/property_formertype"
	propertygenerator "github.com/go-fed/activity/streams/impl/activitystreams/property_generator"
	propertyheight "github.com/go-fed/activity/streams/impl/activitystreams/property_height"
	propertyhref "github.com/go-fed/activity/streams/impl/activitystreams/property_href"
	propertyhreflang "github.com/go-fed/activity/streams/impl/activitystreams/property_hreflang"
	propertyicon "github.com/go-fed/activity/streams/impl/activitystreams/property_icon"
	propertyimage "github.com/go-fed/activity/streams/impl/activitystreams/property_image"
	propertyinbox "github.com/go-fed/activity/streams/impl/activitystreams/property_inbox"
	propertyinreplyto "github.com/go-fed/activity/streams/impl/activitystreams/property_inreplyto"
	propertyinstrument "github.com/go-fed/activity/streams/impl/activitystreams/property_instrument"
	propertyitems "github.com/go-fed/activity/streams/impl/activitystreams/property_items"
	propertylast "github.com/go-fed/activity/streams/impl/activitystreams/property_last"
	propertylatitude "github.com/go-fed/activity/streams/impl/activitystreams/property_latitude"
	propertyliked "github.com/go-fed/activity/streams/impl/activitystreams/property_liked"
	propertylikes "github.com/go-fed/activity/streams/impl/activitystreams/property_likes"
	propertylocation "github.com/go-fed/activity/streams/impl/activitystreams/property_location"
	propertylongitude "github.com/go-fed/activity/streams/impl/activitystreams/property_longitude"
	propertymediatype "github.com/go-fed/activity/streams/impl/activitystreams/property_mediatype"
	propertyname "github.com/go-fed/activity/streams/impl/activitystreams/property_name"
	propertynext "github.com/go-fed/activity/streams/impl/activitystreams/property_next"
	propertyobject "github.com/go-fed/activity/streams/impl/activitystreams/property_object"
	propertyoneof "github.com/go-fed/activity/streams/impl/activitystreams/property_oneof"
	propertyordereditems "github.com/go-fed/activity/streams/impl/activitystreams/property_ordereditems"
	propertyorigin "github.com/go-fed/activity/streams/impl/activitystreams/property_origin"
	propertyoutbox "github.com/go-fed/activity/streams/impl/activitystreams/property_outbox"
	propertypartof "github.com/go-fed/activity/streams/impl/activitystreams/property_partof"
	propertypreferredusername "github.com/go-fed/activity/streams/impl/activitystreams/property_preferredusername"
	propertyprev "github.com/go-fed/activity/streams/impl/activitystreams/property_prev"
	propertypreview "github.com/go-fed/activity/streams/impl/activitystreams/property_preview"
	propertypublished "github.com/go-fed/activity/streams/impl/activitystreams/property_published"
	propertyradius "github.com/go-fed/activity/streams/impl/activitystreams/property_radius"
	propertyrel "github.com/go-fed/activity/streams/impl/activitystreams/property_rel"
	propertyrelationship "github.com/go-fed/activity/streams/impl/activitystreams/property_relationship"
	propertyreplies "github.com/go-fed/activity/streams/impl/activitystreams/property_replies"
	propertyresult "github.com/go-fed/activity/streams/impl/activitystreams/property_result"
	propertyshares "github.com/go-fed/activity/streams/impl/activitystreams/property_shares"
	propertystartindex "github.com/go-fed/activity/streams/impl/activitystreams/property_startindex"
	propertystarttime "github.com/go-fed/activity/streams/impl/activitystreams/property_starttime"
	propertystreams "github.com/go-fed/activity/streams/impl/activitystreams/property_streams"
	propertysubject "github.com/go-fed/activity/streams/impl/activitystreams/property_subject"
	propertysummary "github.com/go-fed/activity/streams/impl/activitystreams/property_summary"
	propertytag "github.com/go-fed/activity/streams/impl/activitystreams/property_tag"
	propertytarget "github.com/go-fed/activity/streams/impl/activitystreams/property_target"
	propertyto "github.com/go-fed/activity/streams/impl/activitystreams/property_to"
	propertytotalitems "github.com/go-fed/activity/streams/impl/activitystreams/property_totalitems"
	propertyunits "github.com/go-fed/activity/streams/impl/activitystreams/property_units"
	propertyupdated "github.com/go-fed/activity/streams/impl/activitystreams/property_updated"
	propertyurl "github.com/go-fed/activity/streams/impl/activitystreams/property_url"
	propertywidth "github.com/go-fed/activity/streams/impl/activitystreams/property_width"
	typeaccept "github.com/go-fed/activity/streams/impl/activitystreams/type_accept"
	typeactivity "github.com/go-fed/activity/streams/impl/activitystreams/type_activity"
	typeadd "github.com/go-fed/activity/streams/impl/activitystreams/type_add"
	typeannounce "github.com/go-fed/activity/streams/impl/activitystreams/type_announce"
	typeapplication "github.com/go-fed/activity/streams/impl/activitystreams/type_application"
	typearrive "github.com/go-fed/activity/streams/impl/activitystreams/type_arrive"
	typearticle "github.com/go-fed/activity/streams/impl/activitystreams/type_article"
	typeaudio "github.com/go-fed/activity/streams/impl/activitystreams/type_audio"
	typeblock "github.com/go-fed/activity/streams/impl/activitystreams/type_block"
	typecollection "github.com/go-fed/activity/streams/impl/activitystreams/type_collection"
	typecollectionpage "github.com/go-fed/activity/streams/impl/activitystreams/type_collectionpage"
	typecreate "github.com/go-fed/activity/streams/impl/activitystreams/type_create"
	typedelete "github.com/go-fed/activity/streams/impl/activitystreams/type_delete"
	typedislike "github.com/go-fed/activity/streams/impl/activitystreams/type_dislike"
	typedocument "github.com/go-fed/activity/streams/impl/activitystreams/type_document"
	typeevent "github.com/go-fed/activity/streams/impl/activitystreams/type_event"
	typeflag "github.com/go-fed/activity/streams/impl/activitystreams/type_flag"
	typefollow "github.com/go-fed/activity/streams/impl/activitystreams/type_follow"
	typegroup "github.com/go-fed/activity/streams/impl/activitystreams/type_group"
	typeignore "github.com/go-fed/activity/streams/impl/activitystreams/type_ignore"
	typeimage "github.com/go-fed/activity/streams/impl/activitystreams/type_image"
	typeintransitiveactivity "github.com/go-fed/activity/streams/impl/activitystreams/type_intransitiveactivity"
	typeinvite "github.com/go-fed/activity/streams/impl/activitystreams/type_invite"
	typejoin "github.com/go-fed/activity/streams/impl/activitystreams/type_join"
	typeleave "github.com/go-fed/activity/streams/impl/activitystreams/type_leave"
	typelike "github.com/go-fed/activity/streams/impl/activitystreams/type_like"
	typelink "github.com/go-fed/activity/streams/impl/activitystreams/type_link"
	typelisten "github.com/go-fed/activity/streams/impl/activitystreams/type_listen"
	typemention "github.com/go-fed/activity/streams/impl/activitystreams/type_mention"
	typemove "github.com/go-fed/activity/streams/impl/activitystreams/type_move"
	typenote "github.com/go-fed/activity/streams/impl/activitystreams/type_note"
	typeobject "github.com/go-fed/activity/streams/impl/activitystreams/type_object"
	typeoffer "github.com/go-fed/activity/streams/impl/activitystreams/type_offer"
	typeorderedcollection "github.com/go-fed/activity/streams/impl/activitystreams/type_orderedcollection"
	typeorderedcollectionpage "github.com/go-fed/activity/streams/impl/activitystreams/type_orderedcollectionpage"
	typeorganization "github.com/go-fed/activity/streams/impl/activitystreams/type_organization"
	typepage "github.com/go-fed/activity/streams/impl/activitystreams/type_page"
	typeperson "github.com/go-fed/activity/streams/impl/activitystreams/type_person"
	typeplace "github.com/go-fed/activity/streams/impl/activitystreams/type_place"
	typeprofile "github.com/go-fed/activity/streams/impl/activitystreams/type_profile"
	typequestion "github.com/go-fed/activity/streams/impl/activitystreams/type_question"
	typeread "github.com/go-fed/activity/streams/impl/activitystreams/type_read"
	typereject "github.com/go-fed/activity/streams/impl/activitystreams/type_reject"
	typerelationship "github.com/go-fed/activity/streams/impl/activitystreams/type_relationship"
	typeremove "github.com/go-fed/activity/streams/impl/activitystreams/type_remove"
	typeservice "github.com/go-fed/activity/streams/impl/activitystreams/type_service"
	typetentativeaccept "github.com/go-fed/activity/streams/impl/activitystreams/type_tentativeaccept"
	typetentativereject "github.com/go-fed/activity/streams/impl/activitystreams/type_tentativereject"
	typetombstone "github.com/go-fed/activity/streams/impl/activitystreams/type_tombstone"
	typetravel "github.com/go-fed/activity/streams/impl/activitystreams/type_travel"
	typeundo "github.com/go-fed/activity/streams/impl/activitystreams/type_undo"
	typeupdate "github.com/go-fed/activity/streams/impl/activitystreams/type_update"
	typevideo "github.com/go-fed/activity/streams/impl/activitystreams/type_video"
	typeview "github.com/go-fed/activity/streams/impl/activitystreams/type_view"
	propertyid "github.com/go-fed/activity/streams/impl/jsonld/property_id"
	propertytype "github.com/go-fed/activity/streams/impl/jsonld/property_type"
	propertyowner "github.com/go-fed/activity/streams/impl/w3idsecurityv1/property_owner"
	propertypublickey "github.com/go-fed/activity/streams/impl/w3idsecurityv1/property_publickey"
	propertypublickeypem "github.com/go-fed/activity/streams/impl/w3idsecurityv1/property_publickeypem"
	typepublickey "github.com/go-fed/activity/streams/impl/w3idsecurityv1/type_publickey"
	vocab "github.com/go-fed/activity/streams/vocab"
)

// Manager manages interface types and deserializations for use by generated code.
// Application code implicitly uses this manager at run-time to create
// concrete implementations of the interfaces.
type Manager struct {
}

// DeserializeAcceptActivityStreams returns the deserialization method for the
// "ActivityStreamsAccept" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeAcceptActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAccept, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsAccept, error) {
		i, err := typeaccept.DeserializeAccept(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeAccuracyPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsAccuracyProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeAccuracyPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAccuracyProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsAccuracyProperty, error) {
		i, err := propertyaccuracy.DeserializeAccuracyProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeActivityActivityStreams returns the deserialization method for the
// "ActivityStreamsActivity" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeActivityActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsActivity, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsActivity, error) {
		i, err := typeactivity.DeserializeActivity(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeActorPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsActorProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeActorPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsActorProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsActorProperty, error) {
		i, err := propertyactor.DeserializeActorProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeAddActivityStreams returns the deserialization method for the
// "ActivityStreamsAdd" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeAddActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAdd, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsAdd, error) {
		i, err := typeadd.DeserializeAdd(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeAltitudePropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsAltitudeProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeAltitudePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAltitudeProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsAltitudeProperty, error) {
		i, err := propertyaltitude.DeserializeAltitudeProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeAnnounceActivityStreams returns the deserialization method for the
// "ActivityStreamsAnnounce" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeAnnounceActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAnnounce, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsAnnounce, error) {
		i, err := typeannounce.DeserializeAnnounce(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeAnyOfPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsAnyOfProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeAnyOfPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAnyOfProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsAnyOfProperty, error) {
		i, err := propertyanyof.DeserializeAnyOfProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeApplicationActivityStreams returns the deserialization method for
// the "ActivityStreamsApplication" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeApplicationActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsApplication, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsApplication, error) {
		i, err := typeapplication.DeserializeApplication(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeArriveActivityStreams returns the deserialization method for the
// "ActivityStreamsArrive" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeArriveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsArrive, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsArrive, error) {
		i, err := typearrive.DeserializeArrive(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeArticleActivityStreams returns the deserialization method for the
// "ActivityStreamsArticle" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeArticleActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsArticle, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsArticle, error) {
		i, err := typearticle.DeserializeArticle(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeAttachmentPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsAttachmentProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeAttachmentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAttachmentProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsAttachmentProperty, error) {
		i, err := propertyattachment.DeserializeAttachmentProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeAttributedToPropertyActivityStreams returns the deserialization
// method for the "ActivityStreamsAttributedToProperty" non-functional
// property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeAttributedToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAttributedToProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsAttributedToProperty, error) {
		i, err := propertyattributedto.DeserializeAttributedToProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeAudiencePropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsAudienceProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeAudiencePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAudienceProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsAudienceProperty, error) {
		i, err := propertyaudience.DeserializeAudienceProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeAudioActivityStreams returns the deserialization method for the
// "ActivityStreamsAudio" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeAudioActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAudio, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsAudio, error) {
		i, err := typeaudio.DeserializeAudio(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeBccPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsBccProperty" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeBccPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsBccProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsBccProperty, error) {
		i, err := propertybcc.DeserializeBccProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeBlockActivityStreams returns the deserialization method for the
// "ActivityStreamsBlock" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeBlockActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsBlock, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsBlock, error) {
		i, err := typeblock.DeserializeBlock(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeBtoPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsBtoProperty" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeBtoPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsBtoProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsBtoProperty, error) {
		i, err := propertybto.DeserializeBtoProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeCcPropertyActivityStreams returns the deserialization method for the
// "ActivityStreamsCcProperty" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeCcPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsCcProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsCcProperty, error) {
		i, err := propertycc.DeserializeCcProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeClosedPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsClosedProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeClosedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsClosedProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsClosedProperty, error) {
		i, err := propertyclosed.DeserializeClosedProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeCollectionActivityStreams returns the deserialization method for the
// "ActivityStreamsCollection" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsCollection, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsCollection, error) {
		i, err := typecollection.DeserializeCollection(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeCollectionPageActivityStreams returns the deserialization method for
// the "ActivityStreamsCollectionPage" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsCollectionPage, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsCollectionPage, error) {
		i, err := typecollectionpage.DeserializeCollectionPage(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeContentPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsContentProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeContentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsContentProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsContentProperty, error) {
		i, err := propertycontent.DeserializeContentProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeContextPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsContextProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeContextPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsContextProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsContextProperty, error) {
		i, err := propertycontext.DeserializeContextProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeCreateActivityStreams returns the deserialization method for the
// "ActivityStreamsCreate" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeCreateActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsCreate, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsCreate, error) {
		i, err := typecreate.DeserializeCreate(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeCurrentPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsCurrentProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeCurrentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsCurrentProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsCurrentProperty, error) {
		i, err := propertycurrent.DeserializeCurrentProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeDeleteActivityStreams returns the deserialization method for the
// "ActivityStreamsDelete" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeDeleteActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsDelete, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsDelete, error) {
		i, err := typedelete.DeserializeDelete(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeDeletedPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsDeletedProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeDeletedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsDeletedProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsDeletedProperty, error) {
		i, err := propertydeleted.DeserializeDeletedProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeDescribesPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsDescribesProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeDescribesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsDescribesProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsDescribesProperty, error) {
		i, err := propertydescribes.DeserializeDescribesProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeDislikeActivityStreams returns the deserialization method for the
// "ActivityStreamsDislike" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeDislikeActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsDislike, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsDislike, error) {
		i, err := typedislike.DeserializeDislike(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeDocumentActivityStreams returns the deserialization method for the
// "ActivityStreamsDocument" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeDocumentActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsDocument, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsDocument, error) {
		i, err := typedocument.DeserializeDocument(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeDurationPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsDurationProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeDurationPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsDurationProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsDurationProperty, error) {
		i, err := propertyduration.DeserializeDurationProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeEndTimePropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsEndTimeProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeEndTimePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsEndTimeProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsEndTimeProperty, error) {
		i, err := propertyendtime.DeserializeEndTimeProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeEventActivityStreams returns the deserialization method for the
// "ActivityStreamsEvent" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeEventActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsEvent, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsEvent, error) {
		i, err := typeevent.DeserializeEvent(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeFirstPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsFirstProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeFirstPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsFirstProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsFirstProperty, error) {
		i, err := propertyfirst.DeserializeFirstProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeFlagActivityStreams returns the deserialization method for the
// "ActivityStreamsFlag" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeFlagActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsFlag, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsFlag, error) {
		i, err := typeflag.DeserializeFlag(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeFollowActivityStreams returns the deserialization method for the
// "ActivityStreamsFollow" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeFollowActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsFollow, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsFollow, error) {
		i, err := typefollow.DeserializeFollow(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeFollowersPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsFollowersProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeFollowersPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsFollowersProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsFollowersProperty, error) {
		i, err := propertyfollowers.DeserializeFollowersProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeFollowingPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsFollowingProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeFollowingPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsFollowingProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsFollowingProperty, error) {
		i, err := propertyfollowing.DeserializeFollowingProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeFormerTypePropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsFormerTypeProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeFormerTypePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsFormerTypeProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsFormerTypeProperty, error) {
		i, err := propertyformertype.DeserializeFormerTypeProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeGeneratorPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsGeneratorProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeGeneratorPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsGeneratorProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsGeneratorProperty, error) {
		i, err := propertygenerator.DeserializeGeneratorProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeGroupActivityStreams returns the deserialization method for the
// "ActivityStreamsGroup" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeGroupActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsGroup, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsGroup, error) {
		i, err := typegroup.DeserializeGroup(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeHeightPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsHeightProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeHeightPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsHeightProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsHeightProperty, error) {
		i, err := propertyheight.DeserializeHeightProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeHrefPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsHrefProperty" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeHrefPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsHrefProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsHrefProperty, error) {
		i, err := propertyhref.DeserializeHrefProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeHreflangPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsHreflangProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeHreflangPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsHreflangProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsHreflangProperty, error) {
		i, err := propertyhreflang.DeserializeHreflangProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeIconPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsIconProperty" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeIconPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsIconProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsIconProperty, error) {
		i, err := propertyicon.DeserializeIconProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeIdPropertyJSONLD returns the deserialization method for the
// "JSONLDIdProperty" non-functional property in the vocabulary "JSONLD"
func (this Manager) DeserializeIdPropertyJSONLD() func(map[string]interface{}, map[string]string) (vocab.JSONLDIdProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.JSONLDIdProperty, error) {
		i, err := propertyid.DeserializeIdProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeIgnoreActivityStreams returns the deserialization method for the
// "ActivityStreamsIgnore" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeIgnoreActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsIgnore, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsIgnore, error) {
		i, err := typeignore.DeserializeIgnore(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeImageActivityStreams returns the deserialization method for the
// "ActivityStreamsImage" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeImageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsImage, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsImage, error) {
		i, err := typeimage.DeserializeImage(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeImagePropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsImageProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeImagePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsImageProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsImageProperty, error) {
		i, err := propertyimage.DeserializeImageProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeInReplyToPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsInReplyToProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeInReplyToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsInReplyToProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsInReplyToProperty, error) {
		i, err := propertyinreplyto.DeserializeInReplyToProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeInboxPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsInboxProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeInboxPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsInboxProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsInboxProperty, error) {
		i, err := propertyinbox.DeserializeInboxProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeInstrumentPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsInstrumentProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeInstrumentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsInstrumentProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsInstrumentProperty, error) {
		i, err := propertyinstrument.DeserializeInstrumentProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeIntransitiveActivityActivityStreams returns the deserialization
// method for the "ActivityStreamsIntransitiveActivity" non-functional
// property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeIntransitiveActivityActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsIntransitiveActivity, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsIntransitiveActivity, error) {
		i, err := typeintransitiveactivity.DeserializeIntransitiveActivity(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeInviteActivityStreams returns the deserialization method for the
// "ActivityStreamsInvite" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeInviteActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsInvite, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsInvite, error) {
		i, err := typeinvite.DeserializeInvite(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeItemsPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsItemsProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeItemsPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsItemsProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsItemsProperty, error) {
		i, err := propertyitems.DeserializeItemsProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeJoinActivityStreams returns the deserialization method for the
// "ActivityStreamsJoin" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeJoinActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsJoin, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsJoin, error) {
		i, err := typejoin.DeserializeJoin(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeLastPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsLastProperty" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeLastPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLastProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsLastProperty, error) {
		i, err := propertylast.DeserializeLastProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeLatitudePropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsLatitudeProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeLatitudePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLatitudeProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsLatitudeProperty, error) {
		i, err := propertylatitude.DeserializeLatitudeProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeLeaveActivityStreams returns the deserialization method for the
// "ActivityStreamsLeave" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeLeaveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLeave, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsLeave, error) {
		i, err := typeleave.DeserializeLeave(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeLikeActivityStreams returns the deserialization method for the
// "ActivityStreamsLike" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeLikeActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLike, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsLike, error) {
		i, err := typelike.DeserializeLike(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeLikedPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsLikedProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeLikedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLikedProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsLikedProperty, error) {
		i, err := propertyliked.DeserializeLikedProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeLikesPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsLikesProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeLikesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLikesProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsLikesProperty, error) {
		i, err := propertylikes.DeserializeLikesProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeLinkActivityStreams returns the deserialization method for the
// "ActivityStreamsLink" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeLinkActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLink, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsLink, error) {
		i, err := typelink.DeserializeLink(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeListenActivityStreams returns the deserialization method for the
// "ActivityStreamsListen" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeListenActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsListen, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsListen, error) {
		i, err := typelisten.DeserializeListen(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeLocationPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsLocationProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeLocationPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLocationProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsLocationProperty, error) {
		i, err := propertylocation.DeserializeLocationProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeLongitudePropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsLongitudeProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeLongitudePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLongitudeProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsLongitudeProperty, error) {
		i, err := propertylongitude.DeserializeLongitudeProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeMediaTypePropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsMediaTypeProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeMediaTypePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsMediaTypeProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsMediaTypeProperty, error) {
		i, err := propertymediatype.DeserializeMediaTypeProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeMentionActivityStreams returns the deserialization method for the
// "ActivityStreamsMention" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeMentionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsMention, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsMention, error) {
		i, err := typemention.DeserializeMention(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeMoveActivityStreams returns the deserialization method for the
// "ActivityStreamsMove" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeMoveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsMove, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsMove, error) {
		i, err := typemove.DeserializeMove(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeNamePropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsNameProperty" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeNamePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsNameProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsNameProperty, error) {
		i, err := propertyname.DeserializeNameProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeNextPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsNextProperty" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeNextPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsNextProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsNextProperty, error) {
		i, err := propertynext.DeserializeNextProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeNoteActivityStreams returns the deserialization method for the
// "ActivityStreamsNote" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeNoteActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsNote, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsNote, error) {
		i, err := typenote.DeserializeNote(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeObjectActivityStreams returns the deserialization method for the
// "ActivityStreamsObject" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeObjectActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsObject, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsObject, error) {
		i, err := typeobject.DeserializeObject(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeObjectPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsObjectProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeObjectPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsObjectProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsObjectProperty, error) {
		i, err := propertyobject.DeserializeObjectProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeOfferActivityStreams returns the deserialization method for the
// "ActivityStreamsOffer" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeOfferActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOffer, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsOffer, error) {
		i, err := typeoffer.DeserializeOffer(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeOneOfPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsOneOfProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeOneOfPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOneOfProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsOneOfProperty, error) {
		i, err := propertyoneof.DeserializeOneOfProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeOrderedCollectionActivityStreams returns the deserialization method
// for the "ActivityStreamsOrderedCollection" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeOrderedCollectionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOrderedCollection, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsOrderedCollection, error) {
		i, err := typeorderedcollection.DeserializeOrderedCollection(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeOrderedCollectionPageActivityStreams returns the deserialization
// method for the "ActivityStreamsOrderedCollectionPage" non-functional
// property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeOrderedCollectionPageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOrderedCollectionPage, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsOrderedCollectionPage, error) {
		i, err := typeorderedcollectionpage.DeserializeOrderedCollectionPage(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeOrderedItemsPropertyActivityStreams returns the deserialization
// method for the "ActivityStreamsOrderedItemsProperty" non-functional
// property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeOrderedItemsPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOrderedItemsProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsOrderedItemsProperty, error) {
		i, err := propertyordereditems.DeserializeOrderedItemsProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeOrganizationActivityStreams returns the deserialization method for
// the "ActivityStreamsOrganization" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeOrganizationActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOrganization, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsOrganization, error) {
		i, err := typeorganization.DeserializeOrganization(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeOriginPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsOriginProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeOriginPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOriginProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsOriginProperty, error) {
		i, err := propertyorigin.DeserializeOriginProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeOutboxPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsOutboxProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeOutboxPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOutboxProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsOutboxProperty, error) {
		i, err := propertyoutbox.DeserializeOutboxProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeOwnerPropertyW3IDSecurityV1 returns the deserialization method for
// the "W3IDSecurityV1OwnerProperty" non-functional property in the vocabulary
// "W3IDSecurityV1"
func (this Manager) DeserializeOwnerPropertyW3IDSecurityV1() func(map[string]interface{}, map[string]string) (vocab.W3IDSecurityV1OwnerProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.W3IDSecurityV1OwnerProperty, error) {
		i, err := propertyowner.DeserializeOwnerProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializePageActivityStreams returns the deserialization method for the
// "ActivityStreamsPage" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializePageActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPage, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsPage, error) {
		i, err := typepage.DeserializePage(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializePartOfPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsPartOfProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializePartOfPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPartOfProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsPartOfProperty, error) {
		i, err := propertypartof.DeserializePartOfProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializePersonActivityStreams returns the deserialization method for the
// "ActivityStreamsPerson" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializePersonActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPerson, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsPerson, error) {
		i, err := typeperson.DeserializePerson(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializePlaceActivityStreams returns the deserialization method for the
// "ActivityStreamsPlace" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializePlaceActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPlace, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsPlace, error) {
		i, err := typeplace.DeserializePlace(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializePreferredUsernamePropertyActivityStreams returns the deserialization
// method for the "ActivityStreamsPreferredUsernameProperty" non-functional
// property in the vocabulary "ActivityStreams"
func (this Manager) DeserializePreferredUsernamePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPreferredUsernameProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsPreferredUsernameProperty, error) {
		i, err := propertypreferredusername.DeserializePreferredUsernameProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializePrevPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsPrevProperty" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializePrevPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPrevProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsPrevProperty, error) {
		i, err := propertyprev.DeserializePrevProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializePreviewPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsPreviewProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializePreviewPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPreviewProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsPreviewProperty, error) {
		i, err := propertypreview.DeserializePreviewProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeProfileActivityStreams returns the deserialization method for the
// "ActivityStreamsProfile" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeProfileActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsProfile, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsProfile, error) {
		i, err := typeprofile.DeserializeProfile(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializePublicKeyPemPropertyW3IDSecurityV1 returns the deserialization
// method for the "W3IDSecurityV1PublicKeyPemProperty" non-functional property
// in the vocabulary "W3IDSecurityV1"
func (this Manager) DeserializePublicKeyPemPropertyW3IDSecurityV1() func(map[string]interface{}, map[string]string) (vocab.W3IDSecurityV1PublicKeyPemProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.W3IDSecurityV1PublicKeyPemProperty, error) {
		i, err := propertypublickeypem.DeserializePublicKeyPemProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializePublicKeyPropertyW3IDSecurityV1 returns the deserialization method
// for the "W3IDSecurityV1PublicKeyProperty" non-functional property in the
// vocabulary "W3IDSecurityV1"
func (this Manager) DeserializePublicKeyPropertyW3IDSecurityV1() func(map[string]interface{}, map[string]string) (vocab.W3IDSecurityV1PublicKeyProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.W3IDSecurityV1PublicKeyProperty, error) {
		i, err := propertypublickey.DeserializePublicKeyProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializePublicKeyW3IDSecurityV1 returns the deserialization method for the
// "W3IDSecurityV1PublicKey" non-functional property in the vocabulary
// "W3IDSecurityV1"
func (this Manager) DeserializePublicKeyW3IDSecurityV1() func(map[string]interface{}, map[string]string) (vocab.W3IDSecurityV1PublicKey, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.W3IDSecurityV1PublicKey, error) {
		i, err := typepublickey.DeserializePublicKey(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializePublishedPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsPublishedProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializePublishedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPublishedProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsPublishedProperty, error) {
		i, err := propertypublished.DeserializePublishedProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeQuestionActivityStreams returns the deserialization method for the
// "ActivityStreamsQuestion" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeQuestionActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsQuestion, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsQuestion, error) {
		i, err := typequestion.DeserializeQuestion(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeRadiusPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsRadiusProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeRadiusPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsRadiusProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsRadiusProperty, error) {
		i, err := propertyradius.DeserializeRadiusProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeReadActivityStreams returns the deserialization method for the
// "ActivityStreamsRead" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeReadActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsRead, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsRead, error) {
		i, err := typeread.DeserializeRead(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeRejectActivityStreams returns the deserialization method for the
// "ActivityStreamsReject" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeRejectActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsReject, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsReject, error) {
		i, err := typereject.DeserializeReject(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeRelPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsRelProperty" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeRelPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsRelProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsRelProperty, error) {
		i, err := propertyrel.DeserializeRelProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeRelationshipActivityStreams returns the deserialization method for
// the "ActivityStreamsRelationship" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeRelationshipActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsRelationship, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsRelationship, error) {
		i, err := typerelationship.DeserializeRelationship(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeRelationshipPropertyActivityStreams returns the deserialization
// method for the "ActivityStreamsRelationshipProperty" non-functional
// property in the vocabulary "ActivityStreams"
func (this Manager) DeserializeRelationshipPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsRelationshipProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsRelationshipProperty, error) {
		i, err := propertyrelationship.DeserializeRelationshipProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeRemoveActivityStreams returns the deserialization method for the
// "ActivityStreamsRemove" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeRemoveActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsRemove, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsRemove, error) {
		i, err := typeremove.DeserializeRemove(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeRepliesPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsRepliesProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeRepliesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsRepliesProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsRepliesProperty, error) {
		i, err := propertyreplies.DeserializeRepliesProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeResultPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsResultProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeResultPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsResultProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsResultProperty, error) {
		i, err := propertyresult.DeserializeResultProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeServiceActivityStreams returns the deserialization method for the
// "ActivityStreamsService" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeServiceActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsService, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsService, error) {
		i, err := typeservice.DeserializeService(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeSharesPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsSharesProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeSharesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsSharesProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsSharesProperty, error) {
		i, err := propertyshares.DeserializeSharesProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeStartIndexPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsStartIndexProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeStartIndexPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsStartIndexProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsStartIndexProperty, error) {
		i, err := propertystartindex.DeserializeStartIndexProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeStartTimePropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsStartTimeProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeStartTimePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsStartTimeProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsStartTimeProperty, error) {
		i, err := propertystarttime.DeserializeStartTimeProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeStreamsPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsStreamsProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeStreamsPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsStreamsProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsStreamsProperty, error) {
		i, err := propertystreams.DeserializeStreamsProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeSubjectPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsSubjectProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeSubjectPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsSubjectProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsSubjectProperty, error) {
		i, err := propertysubject.DeserializeSubjectProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeSummaryPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsSummaryProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeSummaryPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsSummaryProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsSummaryProperty, error) {
		i, err := propertysummary.DeserializeSummaryProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeTagPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsTagProperty" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeTagPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTagProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsTagProperty, error) {
		i, err := propertytag.DeserializeTagProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeTargetPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsTargetProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeTargetPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTargetProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsTargetProperty, error) {
		i, err := propertytarget.DeserializeTargetProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeTentativeAcceptActivityStreams returns the deserialization method
// for the "ActivityStreamsTentativeAccept" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeTentativeAcceptActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTentativeAccept, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsTentativeAccept, error) {
		i, err := typetentativeaccept.DeserializeTentativeAccept(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeTentativeRejectActivityStreams returns the deserialization method
// for the "ActivityStreamsTentativeReject" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeTentativeRejectActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTentativeReject, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsTentativeReject, error) {
		i, err := typetentativereject.DeserializeTentativeReject(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeToPropertyActivityStreams returns the deserialization method for the
// "ActivityStreamsToProperty" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsToProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsToProperty, error) {
		i, err := propertyto.DeserializeToProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeTombstoneActivityStreams returns the deserialization method for the
// "ActivityStreamsTombstone" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeTombstoneActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTombstone, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsTombstone, error) {
		i, err := typetombstone.DeserializeTombstone(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeTotalItemsPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsTotalItemsProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeTotalItemsPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTotalItemsProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsTotalItemsProperty, error) {
		i, err := propertytotalitems.DeserializeTotalItemsProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeTravelActivityStreams returns the deserialization method for the
// "ActivityStreamsTravel" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeTravelActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTravel, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsTravel, error) {
		i, err := typetravel.DeserializeTravel(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeTypePropertyJSONLD returns the deserialization method for the
// "JSONLDTypeProperty" non-functional property in the vocabulary "JSONLD"
func (this Manager) DeserializeTypePropertyJSONLD() func(map[string]interface{}, map[string]string) (vocab.JSONLDTypeProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.JSONLDTypeProperty, error) {
		i, err := propertytype.DeserializeTypeProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeUndoActivityStreams returns the deserialization method for the
// "ActivityStreamsUndo" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeUndoActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsUndo, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsUndo, error) {
		i, err := typeundo.DeserializeUndo(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeUnitsPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsUnitsProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeUnitsPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsUnitsProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsUnitsProperty, error) {
		i, err := propertyunits.DeserializeUnitsProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeUpdateActivityStreams returns the deserialization method for the
// "ActivityStreamsUpdate" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeUpdateActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsUpdate, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsUpdate, error) {
		i, err := typeupdate.DeserializeUpdate(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeUpdatedPropertyActivityStreams returns the deserialization method
// for the "ActivityStreamsUpdatedProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeUpdatedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsUpdatedProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsUpdatedProperty, error) {
		i, err := propertyupdated.DeserializeUpdatedProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeUrlPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsUrlProperty" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeUrlPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsUrlProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsUrlProperty, error) {
		i, err := propertyurl.DeserializeUrlProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeVideoActivityStreams returns the deserialization method for the
// "ActivityStreamsVideo" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeVideoActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsVideo, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsVideo, error) {
		i, err := typevideo.DeserializeVideo(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeViewActivityStreams returns the deserialization method for the
// "ActivityStreamsView" non-functional property in the vocabulary
// "ActivityStreams"
func (this Manager) DeserializeViewActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsView, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsView, error) {
		i, err := typeview.DeserializeView(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}

// DeserializeWidthPropertyActivityStreams returns the deserialization method for
// the "ActivityStreamsWidthProperty" non-functional property in the
// vocabulary "ActivityStreams"
func (this Manager) DeserializeWidthPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsWidthProperty, error) {
	return func(m map[string]interface{}, aliasMap map[string]string) (vocab.ActivityStreamsWidthProperty, error) {
		i, err := propertywidth.DeserializeWidthProperty(m, aliasMap)
		if i == nil {
			return nil, err
		}
		return i, err
	}
}
