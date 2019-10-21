package vocab

// Indicates that the actor accepts the object. The target property can be used in
// certain circumstances to indicate the context into which the object has
// been accepted.
//
// Example 9 (https://www.w3.org/TR/activitystreams-vocabulary/#ex7a-jsonld):
//   {
//     "actor": {
//       "name": "Sally",
//       "type": "Person"
//     },
//     "object": {
//       "actor": "http://john.example.org",
//       "object": {
//         "name": "Going-Away Party for Jim",
//         "type": "Event"
//       },
//       "type": "Invite"
//     },
//     "summary": "Sally accepted an invitation to a party",
//     "type": "Accept"
//   }
//
// Example 10 (https://www.w3.org/TR/activitystreams-vocabulary/#ex7b-jsonld):
//   {
//     "actor": {
//       "name": "Sally",
//       "type": "Person"
//     },
//     "object": {
//       "name": "Joe",
//       "type": "Person"
//     },
//     "summary": "Sally accepted Joe into the club",
//     "target": {
//       "name": "The Club",
//       "type": "Group"
//     },
//     "type": "Accept"
//   }
type ActivityStreamsAccept interface {
	// GetActivityStreamsActor returns the "actor" property if it exists, and
	// nil otherwise.
	GetActivityStreamsActor() ActivityStreamsActorProperty
	// GetActivityStreamsAltitude returns the "altitude" property if it
	// exists, and nil otherwise.
	GetActivityStreamsAltitude() ActivityStreamsAltitudeProperty
	// GetActivityStreamsAttachment returns the "attachment" property if it
	// exists, and nil otherwise.
	GetActivityStreamsAttachment() ActivityStreamsAttachmentProperty
	// GetActivityStreamsAttributedTo returns the "attributedTo" property if
	// it exists, and nil otherwise.
	GetActivityStreamsAttributedTo() ActivityStreamsAttributedToProperty
	// GetActivityStreamsAudience returns the "audience" property if it
	// exists, and nil otherwise.
	GetActivityStreamsAudience() ActivityStreamsAudienceProperty
	// GetActivityStreamsBcc returns the "bcc" property if it exists, and nil
	// otherwise.
	GetActivityStreamsBcc() ActivityStreamsBccProperty
	// GetActivityStreamsBto returns the "bto" property if it exists, and nil
	// otherwise.
	GetActivityStreamsBto() ActivityStreamsBtoProperty
	// GetActivityStreamsCc returns the "cc" property if it exists, and nil
	// otherwise.
	GetActivityStreamsCc() ActivityStreamsCcProperty
	// GetActivityStreamsContent returns the "content" property if it exists,
	// and nil otherwise.
	GetActivityStreamsContent() ActivityStreamsContentProperty
	// GetActivityStreamsContext returns the "context" property if it exists,
	// and nil otherwise.
	GetActivityStreamsContext() ActivityStreamsContextProperty
	// GetActivityStreamsDuration returns the "duration" property if it
	// exists, and nil otherwise.
	GetActivityStreamsDuration() ActivityStreamsDurationProperty
	// GetActivityStreamsEndTime returns the "endTime" property if it exists,
	// and nil otherwise.
	GetActivityStreamsEndTime() ActivityStreamsEndTimeProperty
	// GetActivityStreamsGenerator returns the "generator" property if it
	// exists, and nil otherwise.
	GetActivityStreamsGenerator() ActivityStreamsGeneratorProperty
	// GetActivityStreamsIcon returns the "icon" property if it exists, and
	// nil otherwise.
	GetActivityStreamsIcon() ActivityStreamsIconProperty
	// GetActivityStreamsImage returns the "image" property if it exists, and
	// nil otherwise.
	GetActivityStreamsImage() ActivityStreamsImageProperty
	// GetActivityStreamsInReplyTo returns the "inReplyTo" property if it
	// exists, and nil otherwise.
	GetActivityStreamsInReplyTo() ActivityStreamsInReplyToProperty
	// GetActivityStreamsInstrument returns the "instrument" property if it
	// exists, and nil otherwise.
	GetActivityStreamsInstrument() ActivityStreamsInstrumentProperty
	// GetActivityStreamsLikes returns the "likes" property if it exists, and
	// nil otherwise.
	GetActivityStreamsLikes() ActivityStreamsLikesProperty
	// GetActivityStreamsLocation returns the "location" property if it
	// exists, and nil otherwise.
	GetActivityStreamsLocation() ActivityStreamsLocationProperty
	// GetActivityStreamsMediaType returns the "mediaType" property if it
	// exists, and nil otherwise.
	GetActivityStreamsMediaType() ActivityStreamsMediaTypeProperty
	// GetActivityStreamsName returns the "name" property if it exists, and
	// nil otherwise.
	GetActivityStreamsName() ActivityStreamsNameProperty
	// GetActivityStreamsObject returns the "object" property if it exists,
	// and nil otherwise.
	GetActivityStreamsObject() ActivityStreamsObjectProperty
	// GetActivityStreamsOrigin returns the "origin" property if it exists,
	// and nil otherwise.
	GetActivityStreamsOrigin() ActivityStreamsOriginProperty
	// GetActivityStreamsPreview returns the "preview" property if it exists,
	// and nil otherwise.
	GetActivityStreamsPreview() ActivityStreamsPreviewProperty
	// GetActivityStreamsPublished returns the "published" property if it
	// exists, and nil otherwise.
	GetActivityStreamsPublished() ActivityStreamsPublishedProperty
	// GetActivityStreamsReplies returns the "replies" property if it exists,
	// and nil otherwise.
	GetActivityStreamsReplies() ActivityStreamsRepliesProperty
	// GetActivityStreamsResult returns the "result" property if it exists,
	// and nil otherwise.
	GetActivityStreamsResult() ActivityStreamsResultProperty
	// GetActivityStreamsShares returns the "shares" property if it exists,
	// and nil otherwise.
	GetActivityStreamsShares() ActivityStreamsSharesProperty
	// GetActivityStreamsStartTime returns the "startTime" property if it
	// exists, and nil otherwise.
	GetActivityStreamsStartTime() ActivityStreamsStartTimeProperty
	// GetActivityStreamsSummary returns the "summary" property if it exists,
	// and nil otherwise.
	GetActivityStreamsSummary() ActivityStreamsSummaryProperty
	// GetActivityStreamsTag returns the "tag" property if it exists, and nil
	// otherwise.
	GetActivityStreamsTag() ActivityStreamsTagProperty
	// GetActivityStreamsTarget returns the "target" property if it exists,
	// and nil otherwise.
	GetActivityStreamsTarget() ActivityStreamsTargetProperty
	// GetActivityStreamsTo returns the "to" property if it exists, and nil
	// otherwise.
	GetActivityStreamsTo() ActivityStreamsToProperty
	// GetActivityStreamsUpdated returns the "updated" property if it exists,
	// and nil otherwise.
	GetActivityStreamsUpdated() ActivityStreamsUpdatedProperty
	// GetActivityStreamsUrl returns the "url" property if it exists, and nil
	// otherwise.
	GetActivityStreamsUrl() ActivityStreamsUrlProperty
	// GetJSONLDId returns the "id" property if it exists, and nil otherwise.
	GetJSONLDId() JSONLDIdProperty
	// GetJSONLDType returns the "type" property if it exists, and nil
	// otherwise.
	GetJSONLDType() JSONLDTypeProperty
	// GetTypeName returns the name of this type.
	GetTypeName() string
	// GetUnknownProperties returns the unknown properties for the Accept
	// type. Note that this should not be used by app developers. It is
	// only used to help determine which implementation is LessThan the
	// other. Developers who are creating a different implementation of
	// this type's interface can use this method in their LessThan
	// implementation, but routine ActivityPub applications should not use
	// this to bypass the code generation tool.
	GetUnknownProperties() map[string]interface{}
	// IsExtending returns true if the Accept type extends from the other type.
	IsExtending(other Type) bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this type and the specific properties that are set. The value
	// in the map is the alias used to import the type and its properties.
	JSONLDContext() map[string]string
	// LessThan computes if this Accept is lesser, with an arbitrary but
	// stable determination.
	LessThan(o ActivityStreamsAccept) bool
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format.
	Serialize() (map[string]interface{}, error)
	// SetActivityStreamsActor sets the "actor" property.
	SetActivityStreamsActor(i ActivityStreamsActorProperty)
	// SetActivityStreamsAltitude sets the "altitude" property.
	SetActivityStreamsAltitude(i ActivityStreamsAltitudeProperty)
	// SetActivityStreamsAttachment sets the "attachment" property.
	SetActivityStreamsAttachment(i ActivityStreamsAttachmentProperty)
	// SetActivityStreamsAttributedTo sets the "attributedTo" property.
	SetActivityStreamsAttributedTo(i ActivityStreamsAttributedToProperty)
	// SetActivityStreamsAudience sets the "audience" property.
	SetActivityStreamsAudience(i ActivityStreamsAudienceProperty)
	// SetActivityStreamsBcc sets the "bcc" property.
	SetActivityStreamsBcc(i ActivityStreamsBccProperty)
	// SetActivityStreamsBto sets the "bto" property.
	SetActivityStreamsBto(i ActivityStreamsBtoProperty)
	// SetActivityStreamsCc sets the "cc" property.
	SetActivityStreamsCc(i ActivityStreamsCcProperty)
	// SetActivityStreamsContent sets the "content" property.
	SetActivityStreamsContent(i ActivityStreamsContentProperty)
	// SetActivityStreamsContext sets the "context" property.
	SetActivityStreamsContext(i ActivityStreamsContextProperty)
	// SetActivityStreamsDuration sets the "duration" property.
	SetActivityStreamsDuration(i ActivityStreamsDurationProperty)
	// SetActivityStreamsEndTime sets the "endTime" property.
	SetActivityStreamsEndTime(i ActivityStreamsEndTimeProperty)
	// SetActivityStreamsGenerator sets the "generator" property.
	SetActivityStreamsGenerator(i ActivityStreamsGeneratorProperty)
	// SetActivityStreamsIcon sets the "icon" property.
	SetActivityStreamsIcon(i ActivityStreamsIconProperty)
	// SetActivityStreamsImage sets the "image" property.
	SetActivityStreamsImage(i ActivityStreamsImageProperty)
	// SetActivityStreamsInReplyTo sets the "inReplyTo" property.
	SetActivityStreamsInReplyTo(i ActivityStreamsInReplyToProperty)
	// SetActivityStreamsInstrument sets the "instrument" property.
	SetActivityStreamsInstrument(i ActivityStreamsInstrumentProperty)
	// SetActivityStreamsLikes sets the "likes" property.
	SetActivityStreamsLikes(i ActivityStreamsLikesProperty)
	// SetActivityStreamsLocation sets the "location" property.
	SetActivityStreamsLocation(i ActivityStreamsLocationProperty)
	// SetActivityStreamsMediaType sets the "mediaType" property.
	SetActivityStreamsMediaType(i ActivityStreamsMediaTypeProperty)
	// SetActivityStreamsName sets the "name" property.
	SetActivityStreamsName(i ActivityStreamsNameProperty)
	// SetActivityStreamsObject sets the "object" property.
	SetActivityStreamsObject(i ActivityStreamsObjectProperty)
	// SetActivityStreamsOrigin sets the "origin" property.
	SetActivityStreamsOrigin(i ActivityStreamsOriginProperty)
	// SetActivityStreamsPreview sets the "preview" property.
	SetActivityStreamsPreview(i ActivityStreamsPreviewProperty)
	// SetActivityStreamsPublished sets the "published" property.
	SetActivityStreamsPublished(i ActivityStreamsPublishedProperty)
	// SetActivityStreamsReplies sets the "replies" property.
	SetActivityStreamsReplies(i ActivityStreamsRepliesProperty)
	// SetActivityStreamsResult sets the "result" property.
	SetActivityStreamsResult(i ActivityStreamsResultProperty)
	// SetActivityStreamsShares sets the "shares" property.
	SetActivityStreamsShares(i ActivityStreamsSharesProperty)
	// SetActivityStreamsStartTime sets the "startTime" property.
	SetActivityStreamsStartTime(i ActivityStreamsStartTimeProperty)
	// SetActivityStreamsSummary sets the "summary" property.
	SetActivityStreamsSummary(i ActivityStreamsSummaryProperty)
	// SetActivityStreamsTag sets the "tag" property.
	SetActivityStreamsTag(i ActivityStreamsTagProperty)
	// SetActivityStreamsTarget sets the "target" property.
	SetActivityStreamsTarget(i ActivityStreamsTargetProperty)
	// SetActivityStreamsTo sets the "to" property.
	SetActivityStreamsTo(i ActivityStreamsToProperty)
	// SetActivityStreamsUpdated sets the "updated" property.
	SetActivityStreamsUpdated(i ActivityStreamsUpdatedProperty)
	// SetActivityStreamsUrl sets the "url" property.
	SetActivityStreamsUrl(i ActivityStreamsUrlProperty)
	// SetJSONLDId sets the "id" property.
	SetJSONLDId(i JSONLDIdProperty)
	// SetJSONLDType sets the "type" property.
	SetJSONLDType(i JSONLDTypeProperty)
	// VocabularyURI returns the vocabulary's URI as a string.
	VocabularyURI() string
}
