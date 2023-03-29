CREATE TABLE [dbo].[events](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[name] [nvarchar](255) NOT NULL,
	[date] [date] NOT NULL
) ON [PRIMARY]
GO


CREATE TABLE [dbo].[event_participants](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[user_id] [int] NULL,
	[event_id] [int] NULL
) ON [PRIMARY]
GO

CREATE TABLE [dbo].[invitations](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[meeting_id] [int] NULL,
	[invitee_id] [int] NULL,
	[status] [nvarchar](50) NOT NULL
) ON [PRIMARY]
GO

CREATE TABLE [dbo].[meetings](
    [id] [int] IDENTITY(1,1) NOT NULL,
    [event_id] [int] NULL,
    [organizer_id] [int] NULL,
    [datetime] [datetime] NOT NULL,
    [scheduled] [bit] DEFAULT 0
) ON [PRIMARY];
GO

CREATE TABLE [dbo].[organizations] (
    [id]   INT            IDENTITY (1, 1) NOT NULL,
    [name] NVARCHAR (255) NOT NULL,
    PRIMARY KEY CLUSTERED ([id] ASC)
);

CREATE TABLE [dbo].[users](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[name] [nvarchar](255) NOT NULL,
	[organization_id] [int] NULL
) ON [PRIMARY]
GO